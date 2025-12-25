import { vehicleService, Vehicle } from '@/services/vehicle';
import { gameEvents, GAME_EVENTS } from './EventBus';

export interface BastionState {
  vehicles: Vehicle[];
  selectedVehicleId: string | null;
  isLoading: boolean;
  error: string | null;
}

class BastionSystem {
  private state: BastionState = {
    vehicles: [],
    selectedVehicleId: null,
    isLoading: false,
    error: null,
  };

  getState() {
    return { ...this.state };
  }

  async refreshVehicles(characterId?: string) {
    this.state.isLoading = true;
    this.state.error = null;
    gameEvents.emit(GAME_EVENTS.BASTION_UPDATED, this.getState());

    try {
      const vehicles = await vehicleService.getVehicles(characterId);
      this.state.vehicles = vehicles || [];
      if (this.state.vehicles.length > 0 && !this.state.selectedVehicleId) {
        this.state.selectedVehicleId = this.state.vehicles[0].id;
      }
      gameEvents.emit(GAME_EVENTS.BASTION_UPDATED, this.getState());
    } catch (error: any) {
      this.state.error = error.message;
      gameEvents.emit(GAME_EVENTS.BASTION_UPDATED, this.getState());
    } finally {
      this.state.isLoading = false;
      gameEvents.emit(GAME_EVENTS.BASTION_UPDATED, this.getState());
    }
  }

  selectVehicle(vehicleId: string) {
    this.state.selectedVehicleId = vehicleId;
    gameEvents.emit(GAME_EVENTS.BASTION_UPDATED, this.getState());
  }
}

export const bastionSystem = new BastionSystem();
