type Handler = (data?: any) => void;

class EventBus {
  private events: Record<string, Handler[]> = {};

  on(event: string, handler: Handler) {
    if (!this.events[event]) this.events[event] = [];
    this.events[event].push(handler);
    return () => this.off(event, handler);
  }

  off(event: string, handler: Handler) {
    if (!this.events[event]) return;
    this.events[event] = this.events[event].filter(h => h !== handler);
  }

  emit(event: string, data?: any) {
    if (!this.events[event]) return;
    this.events[event].forEach(handler => handler(data));
  }
}

export const gameEvents = new EventBus();

// Common Event Constants
export const GAME_EVENTS = {
  // Exploration
  EXPLORATION_STARTED: 'EXPLORATION_STARTED',
  EXPLORATION_UPDATED: 'EXPLORATION_UPDATED',
  EXPLORATION_ENDED: 'EXPLORATION_ENDED',
  
  // Combat
  COMBAT_STARTED: 'COMBAT_STARTED',
  COMBAT_UPDATED: 'COMBAT_UPDATED',
  COMBAT_ENDED: 'COMBAT_ENDED',
  
  // Bastion & Gacha
  BASTION_UPDATED: 'BASTION_UPDATED',
  HANGAR_UPDATED: 'HANGAR_UPDATED',
  GACHA_PULLED: 'GACHA_PULLED',
  UNIVERSE_UPDATED: 'UNIVERSE_UPDATED',

  // UI & Effects
  NOTIFICATION: 'NOTIFICATION',
  SCREEN_SHAKE: 'SCREEN_SHAKE',
  PLAYER_DAMAGED: 'PLAYER_DAMAGED',
  ENEMY_DAMAGED: 'ENEMY_DAMAGED',
  ITEM_COLLECTED: 'ITEM_COLLECTED',
};
