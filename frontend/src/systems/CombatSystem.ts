import { combatService, BattleResponse, DamageType, UnitStats } from '@/services/combat';
import { gameEvents, GAME_EVENTS } from './EventBus';

export interface CombatState {
  attackerStats: UnitStats | null;
  defenderStats: UnitStats | null;
  combatLog: string[];
  isProcessing: boolean;
  turn: number;
  isVictory: boolean;
  isDefeat: boolean;
}

class CombatSystem {
  private state: CombatState = {
    attackerStats: null,
    defenderStats: null,
    combatLog: [],
    isProcessing: false,
    turn: 1,
    isVictory: false,
    isDefeat: false,
  };

  getState() {
    return { ...this.state };
  }

  async startCombat(attackerId: string, enemyId: string) {
    this.state = {
      attackerStats: null,
      defenderStats: null,
      combatLog: ['COMBAT ENGAGED'],
      isProcessing: false,
      turn: 1,
      isVictory: false,
      isDefeat: false,
    };
    gameEvents.emit(GAME_EVENTS.COMBAT_STARTED, { attackerId, enemyId });
  }

  async executeAttack(attackerId: string, enemyId: string, type: DamageType) {
    if (this.state.isProcessing || this.state.isVictory || this.state.isDefeat) return;

    this.state.isProcessing = true;
    gameEvents.emit(GAME_EVENTS.COMBAT_UPDATED, this.getState());

    try {
      const data: BattleResponse = await combatService.simulateAttack(attackerId, enemyId, type);
      
      this.state.attackerStats = data.attacker_stats;
      this.state.defenderStats = data.defender_stats;

      const result = data.result;
      let message = `Turn ${this.state.turn}: `;
      if (result.is_miss) {
        message += `Attacker missed!`;
      } else {
        message += `Attacker dealt ${result.final_damage} ${type} damage.`;
        if (result.is_critical) message += ` CRITICAL HIT!`;
      }

      this.state.combatLog = [message, ...this.state.combatLog];
      this.state.turn += 1;

      // Check for victory/defeat
      if (this.state.defenderStats.hp <= 0) {
        this.state.isVictory = true;
        this.state.combatLog = [`Enemy has been destroyed!`, ...this.state.combatLog];
        gameEvents.emit(GAME_EVENTS.COMBAT_ENDED, { result: 'VICTORY' });
      } else if (this.state.attackerStats.hp <= 0) {
        this.state.isDefeat = true;
        this.state.combatLog = [`You have been defeated!`, ...this.state.combatLog];
        gameEvents.emit(GAME_EVENTS.COMBAT_ENDED, { result: 'DEFEAT' });
      }

      gameEvents.emit(GAME_EVENTS.COMBAT_UPDATED, this.getState());
    } catch (error) {
      console.error('Combat error:', error);
      this.state.combatLog = [`Error: Failed to execute attack.`, ...this.state.combatLog];
      gameEvents.emit(GAME_EVENTS.COMBAT_UPDATED, this.getState());
    } finally {
      this.state.isProcessing = false;
      gameEvents.emit(GAME_EVENTS.COMBAT_UPDATED, this.getState());
    }
  }
}

export const combatSystem = new CombatSystem();
