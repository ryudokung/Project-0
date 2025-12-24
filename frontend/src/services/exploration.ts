const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

export interface PlanetLocation {
  id: string;
  name: string;
  description: string;
  rewards: string[];
  requirements: string[];
  allowedModes: string[];
  requiresAtmosphere: boolean;
  suitability: {
    pilot: number;
    mech: number;
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
  suitability: {
    pilot: number;
    mech: number;
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
  suitability_pilot: number;
  suitability_mech: number;
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
  suitability_pilot: number;
  suitability_mech: number;
  coordinates_x: number;
  coordinates_y: number;
}

export const explorationService = {
  async getUniverseMap(): Promise<Sector[]> {
    const response = await fetch(`${API_BASE_URL}/exploration/universe-map`);
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
        suitability: {
          pilot: ss.suitability_pilot,
          mech: ss.suitability_mech
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
          suitability: {
            pilot: l.suitability_pilot,
            mech: l.suitability_mech
          },
          coordinates: { x: l.coordinates_x, y: l.coordinates_y }
        }))
      }))
    }));
  },

  async startExploration(userId: string, subSectorId: string, vehicleId: string, planetLocationId?: string) {
    const response = await fetch(`${API_BASE_URL}/exploration/start`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ 
        user_id: userId, 
        sub_sector_id: subSectorId, 
        vehicle_id: vehicleId,
        planet_location_id: planetLocationId
      }),
    });
    if (!response.ok) {
      throw new Error('Failed to start exploration');
    }
    return response.json();
  },

  async advanceTimeline(threadId: string, vehicleId: string) {
    const response = await fetch(`${API_BASE_URL}/exploration/advance`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ thread_id: threadId, vehicle_id: vehicleId }),
    });
    if (!response.ok) {
      throw new Error('Failed to advance timeline');
    }
    return response.json();
  }
};
