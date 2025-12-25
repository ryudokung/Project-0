const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL 
  ? `${process.env.NEXT_PUBLIC_API_URL}/api/v1` 
  : 'http://localhost:8080/api/v1';

export type DamageType = 'KINETIC' | 'ENERGY' | 'EXPLOSIVE';

export interface UnitStats {
  hp: number;
  max_hp: number;
  base_attack: number;
  target_defense: number;
  defense_efficiency: number;
  accuracy: number;
  evasion: number;
  speed: number;
}

export interface CombatResult {
  final_damage: number;
  is_critical: boolean;
  is_miss: boolean;
  applied_effect?: {
    type: string;
    duration: number;
  };
}

export interface BattleResponse {
  attacker_stats: UnitStats;
  defender_stats: UnitStats;
  result: CombatResult;
}

const getAuthHeaders = () => {
  const token = typeof window !== 'undefined' ? localStorage.getItem('project0_token') : null;
  return {
    'Content-Type': 'application/json',
    ...(token ? { 'Authorization': `Bearer ${token}` } : {}),
  };
};

export const combatService = {
  async simulateAttack(attackerMechId: string, defenderMechId: string, damageType: DamageType = 'KINETIC'): Promise<BattleResponse> {
    const response = await fetch(`${API_BASE_URL}/combat/attack`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({
        attacker_mech_id: attackerMechId,
        defender_mech_id: defenderMechId,
        damage_type: damageType,
      }),
    });

    if (!response.ok) {
      const error = await response.text();
      throw new Error(error || 'Failed to simulate attack');
    }

    return response.json();
  }
};
