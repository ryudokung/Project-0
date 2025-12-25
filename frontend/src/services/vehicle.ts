const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL 
  ? `${process.env.NEXT_PUBLIC_API_URL}/api/v1` 
  : 'http://localhost:8080/api/v1';

export interface Vehicle {
  id: string;
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
  }
};
