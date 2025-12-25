import { mechService, Mech } from '@/services/mech';
import { gameEvents, GAME_EVENTS } from './EventBus';

export interface HangarState {
  mechs: Mech[];
  selectedMechId: string | null;
  isLoading: boolean;
  error: string | null;
}

class HangarSystem {
  private state: HangarState = {
    mechs: [],
    selectedMechId: null,
    isLoading: false,
    error: null,
  };

  getState() {
    return { ...this.state };
  }

  async refreshMechs(characterId?: string) {
    this.state.isLoading = true;
    this.state.error = null;
    gameEvents.emit(GAME_EVENTS.HANGAR_UPDATED, this.getState());

    try {
      const mechs = await mechService.getMechs(characterId);
      this.state.mechs = mechs || [];
      if (this.state.mechs.length > 0 && !this.state.selectedMechId) {
        this.state.selectedMechId = this.state.mechs[0].id;
      }
      gameEvents.emit(GAME_EVENTS.HANGAR_UPDATED, this.getState());
    } catch (error: any) {
      this.state.error = error.message;
      gameEvents.emit(GAME_EVENTS.HANGAR_UPDATED, this.getState());
    } finally {
      this.state.isLoading = false;
      gameEvents.emit(GAME_EVENTS.HANGAR_UPDATED, this.getState());
    }
  }

  selectMech(mechId: string) {
    this.state.selectedMechId = mechId;
    gameEvents.emit(GAME_EVENTS.HANGAR_UPDATED, this.getState());
  }
}

export const hangarSystem = new HangarSystem();
