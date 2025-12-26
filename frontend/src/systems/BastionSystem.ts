import { vehicleService, Vehicle, Item } from '@/services/vehicle';
import { explorationService, PilotStats } from '@/services/exploration';
import { gameEvents, GAME_EVENTS } from './EventBus';

export interface BastionState {
  vehicles: Vehicle[];
  items: Item[];
  pilotStats: PilotStats | null;
  selectedVehicleId: string | null;
  selectedVehicleCP: number;
  isLoading: boolean;
  error: string | null;
}

class BastionSystem {
  private state: BastionState = {
    vehicles: [],
    items: [],
    pilotStats: null,
    selectedVehicleId: null,
    selectedVehicleCP: 0,
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
      const [vehicles, items, pilotStats] = await Promise.all([
        vehicleService.getVehicles(characterId),
        vehicleService.getItems(),
        characterId ? explorationService.getPilotStats(characterId) : Promise.resolve(null)
      ]);

      this.state.vehicles = vehicles || [];
      this.state.items = items || [];
      this.state.pilotStats = pilotStats;

      if (this.state.vehicles.length > 0 && !this.state.selectedVehicleId) {
        this.state.selectedVehicleId = this.state.vehicles[0].id;
      }

      if (this.state.selectedVehicleId) {
        this.state.selectedVehicleCP = await vehicleService.getVehicleCP(this.state.selectedVehicleId);
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

  async selectVehicle(vehicleId: string) {
    this.state.selectedVehicleId = vehicleId;
    try {
      this.state.selectedVehicleCP = await vehicleService.getVehicleCP(vehicleId);
    } catch (e) {
      console.error("Failed to fetch CP", e);
    }
    gameEvents.emit(GAME_EVENTS.BASTION_UPDATED, this.getState());
  }

  async equipItem(itemId: string) {
    if (!this.state.selectedVehicleId) return;
    
    try {
      await vehicleService.equipItem(itemId, this.state.selectedVehicleId);
      // Refresh everything
      await this.refreshVehicles();
      gameEvents.emit(GAME_EVENTS.NOTIFICATION, { message: 'MODULE EQUIPPED', type: 'SUCCESS' });
    } catch (error: any) {
      gameEvents.emit(GAME_EVENTS.NOTIFICATION, { message: error.message, type: 'ERROR' });
    }
  }

  async unequipItem(itemId: string) {
    try {
      await vehicleService.unequipItem(itemId);
      // Refresh everything
      await this.refreshVehicles();
      gameEvents.emit(GAME_EVENTS.NOTIFICATION, { message: 'MODULE UNEQUIPPED', type: 'SUCCESS' });
    } catch (error: any) {
      gameEvents.emit(GAME_EVENTS.NOTIFICATION, { message: error.message, type: 'ERROR' });
    }
  }
}

export const bastionSystem = new BastionSystem();
