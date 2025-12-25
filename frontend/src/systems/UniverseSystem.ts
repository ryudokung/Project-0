import { explorationService, Sector } from '@/services/exploration';
import { gameEvents, GAME_EVENTS } from './EventBus';

export interface UniverseState {
  sectors: Sector[];
  isLoading: boolean;
  error: string | null;
}

class UniverseSystem {
  private state: UniverseState = {
    sectors: [],
    isLoading: false,
    error: null,
  };

  getState() {
    return { ...this.state };
  }

  async fetchMap() {
    this.state.isLoading = true;
    this.state.error = null;
    // We don't have a UNIVERSE_UPDATED event yet, but we can use a generic one or add it
    
    try {
      const data = await explorationService.getUniverseMap();
      this.state.sectors = data;
      gameEvents.emit('UNIVERSE_UPDATED', this.getState());
    } catch (error: any) {
      this.state.error = error.message;
      gameEvents.emit('UNIVERSE_UPDATED', this.getState());
    } finally {
      this.state.isLoading = false;
      gameEvents.emit('UNIVERSE_UPDATED', this.getState());
    }
  }
}

export const universeSystem = new UniverseSystem();
