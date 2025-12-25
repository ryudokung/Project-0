import { explorationService, Encounter, Sector, SubSector, PlanetLocation } from '@/services/exploration';
import { gameEvents, GAME_EVENTS } from './EventBus';

export class ExplorationSystem {
  private static instance: ExplorationSystem;
  
  private constructor() {}

  static getInstance() {
    if (!ExplorationSystem.instance) {
      ExplorationSystem.instance = new ExplorationSystem();
    }
    return ExplorationSystem.instance;
  }

  async startMission(userId: string, subSector: SubSector, vehicleId: string, locationId?: string) {
    try {
      const result = await explorationService.startExploration(
        userId,
        subSector.id,
        vehicleId,
        locationId
      );

      const missionData = {
        o2: result.pilot_stats?.current_o2 || 100,
        fuel: result.pilot_stats?.current_fuel || 100,
        encounters: result.encounters,
        currentEncounter: result.encounters[0],
        expeditionId: result.expedition.id,
        title: locationId ? `PLANET // ${subSector.name}` : `SPACE // ${subSector.name}`
      };

      gameEvents.emit(GAME_EVENTS.NOTIFICATION, { message: 'MISSION STARTED', type: 'SUCCESS' });
      return missionData;
    } catch (error) {
      gameEvents.emit(GAME_EVENTS.NOTIFICATION, { message: 'FAILED TO START MISSION', type: 'ERROR' });
      throw error;
    }
  }

  async advance(expeditionId: string, vehicleId: string) {
    try {
      const result = await explorationService.advanceTimeline(expeditionId, vehicleId);
      
      if (result.pilot_stats?.current_o2 <= 0) {
        gameEvents.emit(GAME_EVENTS.NOTIFICATION, { message: 'O2 DEPLETED', type: 'CRITICAL' });
      }

      return {
        encounter: result.encounter,
        stats: result.pilot_stats
      };
    } catch (error) {
      gameEvents.emit(GAME_EVENTS.NOTIFICATION, { message: 'TIMELINE ADVANCE FAILED', type: 'ERROR' });
      throw error;
    }
  }
}

export const explorationSystem = ExplorationSystem.getInstance();
