const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL 
  ? `${process.env.NEXT_PUBLIC_API_URL}/api/v1` 
  : 'http://localhost:8080/api/v1';

export type TerrainType = 'URBAN' | 'ISLANDS' | 'SKY' | 'DESERT' | 'VOID' | 'SPACE';

export interface PlanetLocation {
  id: string;
  name: string;
  description: string;
  rewards: string[];
  requirements: string[];
  allowedModes: string[];
  requiresAtmosphere: boolean;
  terrain: TerrainType;
  detectionThreshold: number;
  suitability: {
    pilot: number;
    vehicle: number;
  };
  coordinates: { x: number; y: number };
}

export interface SubSector {
  id: string;
  type: string;
  name: string;
  description: string;
  rewards: string[];
  requirements: string[];
  allowedModes: string[];
  requiresAtmosphere: boolean;
  terrain: TerrainType;
  detectionThreshold: number;
  suitability: {
    pilot: number;
    vehicle: number;
  };
  coordinates: { x: number; y: number };
  locations?: PlanetLocation[];
}

export interface Sector {
  id: string;
  name: string;
  description: string;
  difficulty: string;
  coordinates: { x: number; y: number };
  color: string;
  subSectors: SubSector[];
}

interface BackendSector {
  id: string;
  name: string;
  description: string;
  difficulty: string;
  coordinates_x: number;
  coordinates_y: number;
  color: string;
  subSectors: BackendSubSector[];
}

interface BackendSubSector {
  id: string;
  type: string;
  name: string;
  description: string;
  rewards: string[];
  requirements: string[];
  allowed_modes: string[];
  requires_atmosphere: boolean;
  terrain: TerrainType;
  detection_threshold: number;
  suitability_pilot: number;
  suitability_vehicle: number;
  coordinates_x: number;
  coordinates_y: number;
  locations?: BackendPlanetLocation[];
}

interface BackendPlanetLocation {
  id: string;
  name: string;
  description: string;
  rewards: string[];
  requirements: string[];
  allowed_modes: string[];
  requires_atmosphere: boolean;
  terrain: TerrainType;
  detection_threshold: number;
  suitability_pilot: number;
  suitability_vehicle: number;
  coordinates_x: number;
  coordinates_y: number;
}

export interface PilotStats {
  user_id: string;
  character_id: string;
  resonance_level: number;
  resonance_exp: number;
  stress: number;
  xp: number;
  rank: number;
  current_o2: number;
  current_fuel: number;
  scrap_metal: number;
  research_data: number;
  updated_at: string;
  metadata?: {
    radar_level?: number;
    lab_level?: number;
    warp_drive_level?: number;
  };
}

export interface StrategicChoice {
  label: string;
  description: string;
  requirements: string[];
  success_chance: number;
  rewards: string[];
  risks: string[];
}

export interface Node {
  id: string;
  expedition_id: string;
  name: string;
  type: 'STANDARD' | 'RESOURCE' | 'COMBAT' | 'ANOMALY' | 'OUTPOST';
  environment_description: string;
  difficulty_multiplier: number;
  position_index: number;
  choices: StrategicChoice[];
  is_resolved: boolean;
  terrain: TerrainType;
  detectionThreshold: number;
}

export interface Encounter {
  id: string;
  type: string;
  title: string;
  description: string;
  visualPrompt: string;
  image: string;
  enemy_id?: string;
  terrain: TerrainType;
  detectionThreshold: number;
}

interface BackendEncounter {
  id: string;
  type: 'COMBAT' | 'RESOURCE' | 'STORY' | 'REST';
  title: string;
  description: string;
  visual_prompt: string;
  enemy_id?: string;
  terrain: TerrainType;
  detection_threshold: number;
}

export interface ExplorationStartResponse {
  expedition: {
    id: string;
    user_id: string;
    sub_sector_id: string;
    planet_location_id: string | null;
    vehicle_id: string;
    title: string;
    description: string;
    goal: string;
  };
  encounters: Encounter[];
  pilot_stats: PilotStats;
}

export interface AdvanceTimelineResponse {
  encounter: Encounter;
  pilot_stats: PilotStats;
}

interface BackendExplorationStartResponse {
  expedition: any;
  encounters: BackendEncounter[];
  pilot_stats: PilotStats;
}

interface BackendAdvanceTimelineResponse {
  encounter: BackendEncounter;
  pilot_stats: PilotStats;
}

const getAuthHeaders = () => {
  const token = typeof window !== 'undefined' ? localStorage.getItem('project0_token') : null;
  return {
    'Content-Type': 'application/json',
    ...(token ? { 'Authorization': `Bearer ${token}` } : {}),
  };
};

export const explorationService = {
  async getUniverseMap(): Promise<Sector[]> {
    const response = await fetch(`${API_BASE_URL}/exploration/universe-map`, {
      headers: getAuthHeaders(),
    });
    if (!response.ok) {
      throw new Error('Failed to fetch universe map');
    }
    const data: BackendSector[] = await response.json();
    
    // Map backend to frontend format
    return data.map(s => ({
      id: s.id,
      name: s.name,
      description: s.description,
      difficulty: s.difficulty as any,
      coordinates: { x: s.coordinates_x, y: s.coordinates_y },
      color: s.color,
      subSectors: s.subSectors.map(ss => ({
        id: ss.id,
        type: ss.type as any,
        name: ss.name,
        description: ss.description,
        rewards: ss.rewards || [],
        requirements: ss.requirements || [],
        allowedModes: ss.allowed_modes as any || [],
        requiresAtmosphere: ss.requires_atmosphere,
        terrain: ss.terrain,
        detectionThreshold: ss.detection_threshold,
        suitability: {
          pilot: ss.suitability_pilot,
          vehicle: ss.suitability_vehicle
        },
        coordinates: { x: ss.coordinates_x, y: ss.coordinates_y },
        locations: ss.locations?.map(l => ({
          id: l.id,
          name: l.name,
          description: l.description,
          rewards: l.rewards || [],
          requirements: l.requirements || [],
          allowedModes: l.allowed_modes as any || [],
          requiresAtmosphere: l.requires_atmosphere,
          terrain: l.terrain,
          detectionThreshold: l.detection_threshold,
          suitability: {
            pilot: l.suitability_pilot,
            vehicle: l.suitability_vehicle
          },
          coordinates: { x: l.coordinates_x, y: l.coordinates_y }
        }))
      }))
    }));
  },

  async startExploration(userId: string, subSectorId: string, vehicleId: string, planetLocationId?: string): Promise<ExplorationStartResponse> {
    const response = await fetch(`${API_BASE_URL}/exploration/start`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ 
        sub_sector_id: subSectorId, 
        vehicle_id: vehicleId,
        planet_location_id: planetLocationId
      }),
    });
    if (!response.ok) {
      const errorText = await response.text();
      console.error('Exploration start failed:', errorText);
      throw new Error(`Failed to start exploration: ${errorText}`);
    }
    const data: BackendExplorationStartResponse = await response.json();
    return {
      ...data,
      encounters: (data.encounters || []).map(e => ({
        id: e.id,
        type: e.type,
        title: e.title,
        description: e.description,
        visualPrompt: e.visual_prompt,
        enemy_id: e.enemy_id,
        terrain: e.terrain,
        detectionThreshold: e.detection_threshold,
        image: 'https://images.unsplash.com/photo-1614728263952-84ea256f9679?q=80&w=1000&auto=format&fit=crop'
      }))
    };
  },

  async advanceTimeline(expeditionId: string, vehicleId: string): Promise<AdvanceTimelineResponse> {
    const response = await fetch(`${API_BASE_URL}/exploration/advance`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ expedition_id: expeditionId, vehicle_id: vehicleId }),
    });
    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(errorText || 'Failed to advance timeline');
    }
    const data: BackendAdvanceTimelineResponse = await response.json();
    return {
      pilot_stats: data.pilot_stats,
      encounter: {
        id: data.encounter.id,
        type: data.encounter.type,
        title: data.encounter.title,
        description: data.encounter.description,
        visualPrompt: data.encounter.visual_prompt,
        enemy_id: data.encounter.enemy_id,
        terrain: data.encounter.terrain,
        detectionThreshold: data.encounter.detection_threshold,
        image: 'https://images.unsplash.com/photo-1550684848-fac1c5b4e853?q=80&w=1000&auto=format&fit=crop'
      }
    };
  },

  async getTimeline(expeditionId: string): Promise<Node[]> {
    const response = await fetch(`${API_BASE_URL}/exploration/timeline?expedition_id=${expeditionId}`, {
      headers: getAuthHeaders(),
    });
    if (!response.ok) throw new Error('Failed to fetch timeline');
    return response.json();
  },

  async resolveChoice(nodeId: string, choice: string): Promise<{ node: Node, pilot_stats: PilotStats }> {
    const response = await fetch(`${API_BASE_URL}/exploration/resolve`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ node_id: nodeId, choice }),
    });
    if (!response.ok) throw new Error('Failed to resolve choice');
    return response.json();
  },

  async resolveNode(nodeId: string): Promise<Node> {
    const response = await fetch(`${API_BASE_URL}/exploration/resolve-node`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ node_id: nodeId }),
    });
    if (!response.ok) throw new Error('Failed to resolve node');
    return response.json();
  },

  async getPilotStats(characterId: string): Promise<PilotStats> {
    const response = await fetch(`${API_BASE_URL}/game/pilot-stats?character_id=${characterId}`, {
      headers: getAuthHeaders(),
    });
    if (!response.ok) throw new Error('Failed to fetch pilot stats');
    return response.json();
  }
};
