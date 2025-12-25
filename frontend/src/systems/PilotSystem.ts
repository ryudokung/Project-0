import { gameEvents, GAME_EVENTS } from './EventBus';

export interface CharacterData {
  name: string;
  gender: string;
  face_index: number;
  hair_index: number;
}

class PilotSystem {
  async createCharacter(data: CharacterData) {
    const token = localStorage.getItem("project0_token");
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/auth/characters/create`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token}`,
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      const err = await response.text();
      throw new Error(err || 'Failed to create character');
    }

    const result = await response.json();
    gameEvents.emit(GAME_EVENTS.NOTIFICATION, { 
      type: 'SUCCESS', 
      message: `Pilot ${data.name} registered successfully.` 
    });
    return result;
  }
}

export const pilotSystem = new PilotSystem();
