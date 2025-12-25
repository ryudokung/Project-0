const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL 
  ? `${process.env.NEXT_PUBLIC_API_URL}/api/v1` 
  : 'http://localhost:8080/api/v1';

const getAuthHeaders = () => {
  const token = typeof window !== 'undefined' ? localStorage.getItem('project0_token') : null;
  return {
    'Content-Type': 'application/json',
    ...(token ? { 'Authorization': `Bearer ${token}` } : {}),
  };
};

export type PullType = 'DAILY_SIGNAL' | 'RELIC_SIGNAL' | 'SINGULARITY_SIGNAL';

export interface GachaPullResponse {
  items: any[];
  pity_relic_count: number;
  pity_singularity_count: number;
  total_pulls: number;
}

export const gachaService = {
  async pull(type: PullType): Promise<GachaPullResponse> {
    const response = await fetch(`${API_BASE_URL}/gacha/pull`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ pull_type: type }),
    });

    if (!response.ok) {
      const error = await response.text();
      throw new Error(error || 'Failed to pull signal');
    }

    return response.json();
  }
};
