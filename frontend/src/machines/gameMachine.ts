import { createMachine, assign } from 'xstate';
import { Sector, SubSector, PlanetLocation, Encounter, Node } from '@/services/exploration';

export type GameEvent = 
  | { type: 'UPDATE_USER'; user: any }
  | { type: 'UPDATE_PILOT_STATS'; stats: any }
  | { type: 'UPDATE_STATS'; o2: number; fuel: number; scrap?: number; research?: number; encounters: Encounter[]; currentEncounter: Encounter | null; expeditionId: string; title: string }
  | { type: 'UPDATE_TIMELINE'; timeline: Node[]; currentNode?: Node }
  | { type: 'RESET_EXPEDITION' }
  | { type: 'START'; user?: any }
  | { type: 'SUCCESS' }
  | { type: 'CHARACTER_CREATED'; character: any }
  | { type: 'DEPLOY' }
  | { type: 'OPEN_GACHA' }
  | { type: 'OPEN_RESEARCH' }
  | { type: 'UPDATE_VEHICLES'; vehicles: any[] }
  | { type: 'BACK' }
  | { type: 'SELECT_SECTOR'; sector: Sector }
  | { type: 'SELECT_SUBSECTOR'; subSector: SubSector }
  | { type: 'SELECT_PLANET_LOCATION'; location: PlanetLocation }
  | { type: 'SELECT_VEHICLE'; vehicle: any }
  | { type: 'CONFIRM_DEPLOYMENT'; isPlanet?: boolean }
  | { type: 'ENTER_COMBAT'; enemyId: string }
  | { type: 'COMBAT_END' }
  | { type: 'MISSION_END' }
  | { type: 'NEXT_NODE' }
  | { type: 'RETURN' }
  | { type: 'RESOLVE_NODE'; node: Node }
  | { type: 'SCAN' }
  | { type: 'AUTH_READY'; user: any };

export interface GameContext {
  user: any | null;
  pilotStats: any | null;
  o2: number;
  fuel: number;
  scrap: number;
  research: number;
  activeEnemyId: string | null;
  encounters: Encounter[];
  currentEncounter: Encounter | null;
  timeline: Node[];
  currentNode: Node | null;
  expeditionTitle: string;
  currentExpeditionId: string | null;
  selectedSector: Sector | null;
  selectedSubSector: SubSector | null;
  selectedPlanetLocation: PlanetLocation | null;
  vehicles: any[];
  selectedVehicle: any | null;
}

export const gameMachine = createMachine({
  id: 'game',
  types: {} as {
    context: GameContext;
    events: GameEvent;
  },
  initial: 'initializing',
  context: {
    user: null,
    pilotStats: null,
    o2: 100,
    fuel: 100,
    scrap: 0,
    research: 0,
    activeEnemyId: null,
    encounters: [],
    currentEncounter: null,
    timeline: [],
    currentNode: null,
    expeditionTitle: 'THE SILENT SIGNAL',
    currentExpeditionId: null,
    selectedSector: null,
    selectedSubSector: null,
    selectedPlanetLocation: null,
    vehicles: [],
    selectedVehicle: null,
  } as GameContext,
  on: {
    UPDATE_USER: {
      actions: assign({ user: ({ event }) => event.user })
    },
    UPDATE_PILOT_STATS: {
      actions: assign({ 
        pilotStats: ({ event }) => event.stats,
        research: ({ event }) => event.stats.research_data,
        scrap: ({ event }) => event.stats.scrap_metal
      })
    },
    UPDATE_STATS: {
      actions: assign({
        o2: ({ event }) => event.o2,
        fuel: ({ event }) => event.fuel,
        scrap: ({ event, context }) => event.scrap !== undefined ? event.scrap : context.scrap,
        research: ({ event, context }) => event.research !== undefined ? event.research : context.research,
        encounters: ({ event }) => event.encounters,
        currentEncounter: ({ event }) => event.currentEncounter,
        currentExpeditionId: ({ event }) => event.expeditionId,
        expeditionTitle: ({ event }) => event.title
      })
    },
    UPDATE_TIMELINE: {
      actions: assign({
        timeline: ({ event }) => event.timeline,
        currentNode: ({ event }) => event.currentNode || event.timeline.find((n: Node) => !n.is_resolved) || event.timeline[0]
      })
    },
    RESET_EXPEDITION: {
      actions: assign({
        timeline: [],
        currentNode: null,
        encounters: [],
        currentEncounter: null,
        currentExpeditionId: null,
        selectedSubSector: null,
        selectedPlanetLocation: null,
      })
    }
  },
  states: {
    initializing: {
      on: {
        AUTH_READY: [
          { 
            target: 'characterCreation', 
            guard: ({ context, event }) => {
              const u = event.user || context.user;
              return u && !u.active_character_id;
            },
            actions: assign({ user: ({ context, event }) => event.user || context.user })
          },
          { 
            target: 'bastion',
            guard: ({ context, event }) => {
              const u = event.user || context.user;
              return u && !!u.active_character_id;
            },
            actions: assign({ user: ({ context, event }) => event.user || context.user })
          },
          { 
            target: 'landing', 
            actions: assign({ user: ({ event }) => event.user })
          }
        ],
        UPDATE_USER: {
          actions: assign({ user: ({ event }) => event.user })
        }
      }
    },
    landing: {
      on: {
        START: [
          { 
            target: 'characterCreation', 
            guard: ({ context, event }) => {
              const u = event.user || context.user;
              return u && !u.active_character_id;
            },
            actions: assign({ user: ({ context, event }) => event.user || context.user })
          },
          { 
            target: 'bastion', 
            guard: ({ context, event }) => {
              const u = event.user || context.user;
              return u && !!u.active_character_id;
            },
            actions: assign({ user: ({ context, event }) => event.user || context.user })
          }
        ],
        UPDATE_USER: [
          { 
            target: 'characterCreation', 
            guard: ({ event }) => event.user && !event.user.active_character_id,
            actions: assign({ user: ({ event }) => event.user })
          },
          { 
            target: 'bastion', 
            guard: ({ event }) => event.user && !!event.user.active_character_id,
            actions: assign({ user: ({ event }) => event.user })
          },
          {
            actions: assign({ user: ({ event }) => event.user })
          }
        ],
        AUTH_READY: [
          { 
            target: 'characterCreation', 
            guard: ({ event }) => event.user && !event.user.active_character_id,
            actions: assign({ user: ({ event }) => event.user })
          },
          { 
            target: 'bastion', 
            guard: ({ event }) => event.user && !!event.user.active_character_id,
            actions: assign({ user: ({ event }) => event.user })
          }
        ]
      }
    },
    characterCreation: {
      on: {
        SUCCESS: 'bastion',
        CHARACTER_CREATED: {
          target: 'bastion',
          actions: assign({ 
            user: ({ context, event }) => ({ 
              ...context.user, 
              active_character_id: event.character.id,
              active_character: event.character
            }) 
          })
        }
      }
    },
    bastion: {
      on: {
        DEPLOY: 'map',
        OPEN_GACHA: 'gacha',
        OPEN_RESEARCH: 'research',
        UPDATE_VEHICLES: {
          actions: assign({ 
            vehicles: ({ event }) => event.vehicles,
            selectedVehicle: ({ context, event }) => context.selectedVehicle || event.vehicles[0]
          })
        }
      }
    },
    gacha: {
      on: {
        BACK: 'bastion'
      }
    },
    research: {
      on: {
        BACK: 'bastion'
      }
    },
    map: {
      on: {
        SELECT_SECTOR: {
          actions: assign({ selectedSector: ({ event }) => event.sector }),
          target: 'locationScan'
        },
        BACK: 'bastion'
      }
    },
    locationScan: {
      on: {
        SELECT_SUBSECTOR: {
          actions: assign({ selectedSubSector: ({ event }) => event.subSector })
        },
        SELECT_VEHICLE: {
          actions: assign({ selectedVehicle: ({ event }) => event.vehicle })
        },
        CONFIRM_DEPLOYMENT: [
          { target: 'planetSurface', guard: ({ event }) => !!event.isPlanet },
          { target: 'exploration' }
        ],
        BACK: 'map'
      }
    },
    planetSurface: {
      on: {
        SELECT_PLANET_LOCATION: {
          actions: assign({ selectedPlanetLocation: ({ event }) => event.location })
        },
        SELECT_VEHICLE: {
          actions: assign({ selectedVehicle: ({ event }) => event.vehicle })
        },
        CONFIRM_DEPLOYMENT: 'exploration',
        BACK: 'locationScan'
      }
    },
    exploration: {
      on: {
        RESOLVE_NODE: {
          actions: assign({
            currentNode: ({ event }) => event.node,
            timeline: ({ context, event }) => context.timeline.map((n: Node) => n.id === event.node.id ? event.node : n)
          })
        },
        ENTER_COMBAT: {
          actions: assign({ activeEnemyId: ({ event }) => event.enemyId }),
          target: 'combat'
        },
        NEXT_NODE: [
          {
            guard: ({ context }) => {
              const currentIndex = context.timeline.findIndex(n => n.id === context.currentNode?.id);
              return currentIndex >= context.timeline.length - 1;
            },
            target: 'debrief'
          },
          {
            actions: assign({
              currentNode: ({ context }) => {
                const currentIndex = context.timeline.findIndex(n => n.id === context.currentNode?.id);
                return context.timeline[currentIndex + 1] || context.currentNode;
              }
            })
          }
        ],
        MISSION_END: 'debrief'
      }
    },
    combat: {
      on: {
        COMBAT_END: 'exploration'
      }
    },
    debrief: {
      on: {
        RETURN: {
          target: 'bastion',
          actions: assign({
            timeline: [],
            currentNode: null,
            encounters: [],
            currentEncounter: null,
            currentExpeditionId: null,
            selectedSubSector: null,
            selectedPlanetLocation: null,
          })
        }
      }
    }
  }
});
