import { createMachine, assign } from 'xstate';
import { Sector, SubSector, PlanetLocation, Encounter, Node } from '@/services/exploration';

export interface GameContext {
  user: any | null;
  o2: number;
  fuel: number;
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
  initial: 'initializing',
  context: {
    user: null,
    o2: 100,
    fuel: 100,
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
        UPDATE_STATS: {
          actions: assign({
            o2: ({ event }) => event.o2,
            fuel: ({ event }) => event.fuel,
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
        RETURN: 'bastion'
      }
    }
  }
});
