import { createMachine, assign } from 'xstate';
import { Sector, SubSector, PlanetLocation, Encounter } from '@/services/exploration';

export interface GameContext {
  user: any | null;
  o2: number;
  fuel: number;
  activeEnemyId: string | null;
  encounters: Encounter[];
  currentEncounter: Encounter | null;
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
    expeditionTitle: 'THE SILENT SIGNAL',
    currentExpeditionId: null,
    selectedSector: null,
    selectedSubSector: null,
    selectedPlanetLocation: null,
    vehicles: [],
    selectedVehicle: null,
  } as GameContext,
  states: {
    initializing: {
      on: {
        AUTH_READY: [
          { target: 'landing', guard: ({ context }) => !context.user },
          { target: 'characterCreation', guard: ({ context }) => context.user && !context.user.active_character_id },
          { target: 'hangar' }
        ],
        UPDATE_USER: {
          actions: assign({ user: ({ event }) => event.user })
        }
      }
    },
    landing: {
      on: {
        START: [
          { target: 'characterCreation', guard: ({ context }) => context.user && !context.user.active_character_id },
          { target: 'hangar', guard: ({ context }) => !!context.user },
        ],
        UPDATE_USER: {
          actions: assign({ user: ({ event }) => event.user })
        }
      }
    },
    characterCreation: {
      on: {
        SUCCESS: 'hangar'
      }
    },
    hangar: {
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
        BACK: 'hangar'
      }
    },
    map: {
      on: {
        SELECT_SECTOR: {
          actions: assign({ selectedSector: ({ event }) => event.sector }),
          target: 'locationScan'
        },
        BACK: 'hangar'
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
        RETURN: 'hangar'
      }
    }
  }
});
