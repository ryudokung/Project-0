import { gachaService, PullType } from '@/services/gacha';
import { gameEvents, GAME_EVENTS } from './EventBus';

export interface GachaState {
  isPulling: boolean;
  results: any[] | null;
  error: string | null;
}

class GachaSystem {
  private state: GachaState = {
    isPulling: false,
    results: null,
    error: null,
  };

  getState() {
    return { ...this.state };
  }

  async pull(type: PullType) {
    this.state.isPulling = true;
    this.state.error = null;
    this.state.results = null;
    gameEvents.emit(GAME_EVENTS.GACHA_PULLED, this.getState());

    try {
      const data = await gachaService.pull(type);
      this.state.results = data.items;
      gameEvents.emit(GAME_EVENTS.GACHA_PULLED, this.getState());
      
      // Notify other systems that inventory might have changed
      gameEvents.emit(GAME_EVENTS.HANGAR_UPDATED); 
    } catch (err: any) {
      this.state.error = err.message;
      gameEvents.emit(GAME_EVENTS.GACHA_PULLED, this.getState());
    } finally {
      this.state.isPulling = false;
      gameEvents.emit(GAME_EVENTS.GACHA_PULLED, this.getState());
    }
  }

  clearResults() {
    this.state.results = null;
    this.state.error = null;
    gameEvents.emit(GAME_EVENTS.GACHA_PULLED, this.getState());
  }
}

export const gachaSystem = new GachaSystem();
