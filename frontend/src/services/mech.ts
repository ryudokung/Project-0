const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL 
  ? `${process.env.NEXT_PUBLIC_API_URL}/api/v1` 
  : 'http://localhost:8080/api/v1';

export interface Mech {
  id: string;
  model: string;
  hp: number;
  attack: number;
  defense: number;
  speed: number;
  character_id: string;
}

const getAuthHeaders = () => {
  const token = typeof window !== 'undefined' ? localStorage.getItem('project0_token') : null;
  return {
    'Content-Type': 'application/json',
    ...(token ? { 'Authorization': `Bearer ${token}` } : {}),
  };
};

export const mechService = {
  async getMechs(characterId?: string): Promise<Mech[]> {
    const url = characterId 
      ? `${API_BASE_URL}/mechs?character_id=${characterId}`
      : `${API_BASE_URL}/mechs`;
      
    const response = await fetch(url, {
      headers: getAuthHeaders(),
    });

    if (!response.ok) {
      throw new Error('Failed to fetch mechs');
    }

    return response.json();
  }
};
