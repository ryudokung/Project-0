const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL 
  ? `${process.env.NEXT_PUBLIC_API_URL}/api/v1` 
  : 'http://localhost:8080/api/v1';

export interface Vehicle {
  id: string;
  name: string;
  class: string;
  vehicle_type: string;
  stats: {
    hp: number;
    attack: number;
    defense: number;
    speed: number;
  };
  character_id?: string;
  rarity: string;
  tier: number;
  cr: number;
  is_void_touched: boolean;
}

export interface Item {
  id: string;
  name: string;
  item_type: string;
  rarity: string;
  tier: number;
  slot?: string;
  damage_type?: string; // KINETIC, ENERGY, VOID
  series_id?: string;   // For Set Synergy
  durability: number;
  max_durability: number;
  condition: string;
  is_nft?: boolean;
  token_id?: string;
  stats: {
    hp?: number;
    attack?: number;
    defense?: number;
    speed?: number;
    bonus_hp?: number;
    bonus_attack?: number;
    bonus_defense?: number;
  };
  is_equipped: boolean;
  parent_item_id?: string;
}

const getAuthHeaders = () => {
  const token = typeof window !== 'undefined' ? localStorage.getItem('project0_token') : null;
  return {
    'Content-Type': 'application/json',
    ...(token ? { 'Authorization': `Bearer ${token}` } : {}),
  };
};

export const vehicleService = {
  async getVehicles(characterId?: string): Promise<Vehicle[]> {
    const url = characterId 
      ? `${API_BASE_URL}/vehicles?character_id=${characterId}`
      : `${API_BASE_URL}/vehicles`;
      
    const response = await fetch(url, {
      headers: getAuthHeaders(),
    });

    if (!response.ok) {
      throw new Error('Failed to fetch vehicles');
    }

    const data = await response.json();
    return data || [];
  },

  async getItems(): Promise<Item[]> {
    const response = await fetch(`${API_BASE_URL}/items`, {
      headers: getAuthHeaders(),
    });
    if (!response.ok) {
      throw new Error('Failed to fetch items');
    }
    return await response.json();
  },

  async getVehicleCP(vehicleId: string): Promise<number> {
    const response = await fetch(`${API_BASE_URL}/vehicles/cp?id=${vehicleId}`, {
      headers: getAuthHeaders(),
    });
    if (!response.ok) {
      throw new Error('Failed to fetch vehicle CP');
    }
    const data = await response.json();
    return data.cp;
  },

  async equipItem(itemId: string, vehicleId: string): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/vehicles/equip`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ item_id: itemId, vehicle_id: vehicleId }),
    });
    if (!response.ok) {
      throw new Error('Failed to equip item');
    }
  },

  async unequipItem(itemId: string): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/vehicles/unequip`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ item_id: itemId }),
    });
    if (!response.ok) {
      throw new Error('Failed to unequip item');
    }
  },

  async repairItem(itemId: string): Promise<{ item: Item; cost: number }> {
    const response = await fetch(`${API_BASE_URL}/vehicles/repair`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ item_id: itemId }),
    });
    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.error || 'Failed to repair item');
    }
    return await response.json();
  },

  async mintItem(itemId: string): Promise<{ status: string; token_id: string }> {
    const response = await fetch(`${API_BASE_URL}/vehicles/mint`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ item_id: itemId }),
    });
    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.error || 'Failed to mint item');
    }
    return await response.json();
  },};
