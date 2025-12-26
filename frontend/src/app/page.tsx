'use client';

import { useState, useEffect } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { useMachine } from '@xstate/react';
import { gameMachine } from '@/machines/gameMachine';
import { Sector, SubSector, PlanetLocation } from '@/services/exploration';
import { explorationSystem } from '@/systems/ExplorationSystem';
import { bastionSystem, BastionState } from '@/systems/BastionSystem';
import { universeSystem, UniverseState } from '@/systems/UniverseSystem';
import { gameEvents, GAME_EVENTS } from '@/systems/EventBus';
import { useAuthSync } from '@/hooks/use-auth-sync';

// Components
import LandingStage from '@/components/game/LandingStage';
import CharacterCreationStage from '@/components/game/CharacterCreationStage';
import Bastion from '@/components/game/Bastion';
import UniverseMap from '@/components/game/UniverseMap';
import LocationScan from '@/components/game/LocationScan';
import PlanetSurface from '@/components/game/PlanetSurface';
import ExplorationLoop from '@/components/game/ExplorationLoop';
import TimelineView from '@/components/game/TimelineView';
import CombatStage from '@/components/game/CombatStage';
import GachaStage from '@/components/game/GachaStage';
import NotificationSystem from '@/components/game/NotificationSystem';

type GameStage = 'LANDING' | 'CHARACTER_CREATION' | 'BASTION' | 'MAP' | 'LOCATION_SCAN' | 'PLANET_SURFACE' | 'EXPLORATION' | 'COMBAT' | 'DEBRIEF' | 'GACHA';
type DeploymentMode = 'PILOT' | 'SPEEDER' | 'MECH' | 'TANK' | 'SHIP' | 'EXOSUIT' | 'HAULER';

interface MothershipUpgrades {
  atmosphericEntry: boolean;
  quantumGate: boolean;
  miningDrill: boolean;
  hackingModule: boolean;
  radarLevel: number;
  scannerLevel: number;
}

export default function UnifiedGamePage() {
  const { user, backendToken, isLoading: authLoading } = useAuthSync();
  const [state, send] = useMachine(gameMachine);
  const { 
    o2, fuel, activeEnemyId, encounters, currentEncounter, 
    expeditionTitle, vehicles, selectedVehicle, selectedSector, 
    selectedSubSector, selectedPlanetLocation, currentExpeditionId,
    timeline, currentNode
  } = state.context;
  
  const [sectors, setSectors] = useState<Sector[]>([]);
  const [isTransitioning, setIsTransitioning] = useState(false);
  
  // Mothership & Inventory
  const [inventory, setInventory] = useState<string[]>(['Basic O2 Tank']);
  const [upgrades, setUpgrades] = useState<MothershipUpgrades>({
    atmosphericEntry: false,
    quantumGate: false,
    miningDrill: false,
    hackingModule: false,
    radarLevel: 1,
    scannerLevel: 1
  });

  // Sync User to Machine
  useEffect(() => {
    if (user) {
      console.log('Updating user in machine:', user);
      send({ type: 'UPDATE_USER', user });
    }
  }, [user, send]);

  // Handle Initial Stage based on Auth
  useEffect(() => {
    if (!authLoading) {
      console.log('Auth ready, sending AUTH_READY:', user);
      send({ type: 'AUTH_READY', user });
    }
  }, [authLoading, user, send]);

  // Log state changes
  useEffect(() => {
    console.log('Current Machine State:', state.value);
    console.log('Current Context User:', state.context.user);
  }, [state.value, state.context.user]);

  // System Listeners
  useEffect(() => {
    const unsubUniverse = gameEvents.on(GAME_EVENTS.UNIVERSE_UPDATED, (state: UniverseState) => {
      setSectors(state.sectors);
    });

    const unsubBastion = gameEvents.on(GAME_EVENTS.BASTION_UPDATED, (state: BastionState) => {
      if (state.vehicles && state.vehicles.length > 0) {
        const mappedVehicles = state.vehicles.map((m: any) => {
          const vehicleClass = m.class || m.model || 'UNKNOWN';
          const vehicleType = m.vehicle_type || (vehicleClass.includes('TANK') ? 'TANK' : vehicleClass.includes('SHIP') ? 'SHIP' : 'VEHICLE');
          
          return {
            id: m.id,
            class: vehicleClass,
            type: vehicleType as DeploymentMode,
            rarity: m.rarity || 'COMMON',
            stats: m.stats || { 
              hp: m.hp || 100, 
              attack: m.attack || 10, 
              defense: m.defense || 10, 
              speed: m.speed || 10 
            }
          };
        });
        send({ type: 'UPDATE_VEHICLES', vehicles: mappedVehicles });
      }
    });

    // Initial Fetch
    universeSystem.fetchMap();
    if (user?.active_character_id || user?.id) {
      bastionSystem.refreshVehicles(user.active_character_id || user.id);
    }

    return () => {
      unsubUniverse();
      unsubBastion();
    };
  }, [user?.id, user?.active_character_id, send]);

  const canDeploy = (target: { allowedModes: string[], requiresAtmosphere: boolean }) => {
    const currentMode = selectedVehicle ? selectedVehicle.type : 'PILOT';
    const modeAllowed = target.allowedModes.includes(currentMode);
    const entryAllowed = !target.requiresAtmosphere || upgrades.atmosphericEntry || upgrades.quantumGate;
    return modeAllowed && entryAllowed;
  };

  const meetsRequirements = (reqs: string[] = []) => {
    return reqs.every(r => {
      if (r === 'Mining Drill') return upgrades.miningDrill;
      if (r === 'Hacking Module') return upgrades.hackingModule;
      return inventory.includes(r);
    });
  };

  const confirmDeployment = async () => {
    if (!selectedSubSector || !user?.id) return;
    
    if (selectedSubSector.type === 'PLANET' && selectedSubSector.locations) {
      send({ type: 'CONFIRM_DEPLOYMENT', isPlanet: true });
    } else {
      try {
        setIsTransitioning(true);
        const result = await explorationSystem.startMission(
          user.id,
          selectedSubSector,
          selectedVehicle?.id || '00000000-0000-0000-0000-000000000000'
        );
        
        // Fetch Timeline
        const timeline = await explorationSystem.getTimeline(result.expeditionId);
        
        send({ 
          type: 'UPDATE_TIMELINE', 
          timeline,
          currentNode: timeline[0]
        });
        send({ type: 'UPDATE_STATS', ...result });
        send({ type: 'CONFIRM_DEPLOYMENT', isPlanet: false });
      } catch (error) {
        console.error('Failed to start exploration:', error);
      } finally {
        setIsTransitioning(false);
      }
    }
  };

  const confirmPlanetDeployment = async () => {
    if (!selectedPlanetLocation || !selectedSubSector || !user?.id) return;
    
    try {
      setIsTransitioning(true);
      const result = await explorationSystem.startMission(
        user.id,
        selectedSubSector,
        selectedVehicle?.id || '00000000-0000-0000-0000-000000000000',
        selectedPlanetLocation.id
      );
      
      // Fetch Timeline
      const timeline = await explorationSystem.getTimeline(result.expeditionId);
      
      send({ type: 'UPDATE_STATS', ...result });
      send({ type: 'UPDATE_TIMELINE', timeline });
      send({ type: 'CONFIRM_DEPLOYMENT' });
    } catch (error) {
      console.error('Failed to start planet exploration:', error);
    } finally {
      setIsTransitioning(false);
    }
  };

  const resolveChoice = async (nodeId: string, choice: string) => {
    if (isTransitioning) return;
    try {
      setIsTransitioning(true);
      const result = await explorationSystem.resolveChoice(nodeId, choice);
      if (result) {
        send({ type: 'RESOLVE_NODE', node: result.node });
        send({ 
          type: 'UPDATE_STATS', 
          o2: result.pilot_stats?.current_o2 ?? o2,
          fuel: result.pilot_stats?.current_fuel ?? fuel,
          encounters,
          currentEncounter,
          expeditionId: currentExpeditionId,
          title: expeditionTitle
        });

        if (result.node.type === 'COMBAT') {
          // For combat nodes, we might want to trigger combat stage
        }
      }
    } catch (error) {
      console.error('Failed to resolve choice:', error);
    } finally {
      setIsTransitioning(false);
    }
  };

  const advanceTimeline = async () => {
    if (isTransitioning || !currentExpeditionId) return;

    // Check if we are at the last node
    const currentIndex = timeline.findIndex(n => n.id === currentNode?.id);
    if (currentIndex === timeline.length - 1) {
      send({ type: 'MISSION_END' });
      return;
    }

    try {
      setIsTransitioning(true);
      const result = await explorationSystem.advance(
        currentExpeditionId,
        selectedVehicle?.id || '00000000-0000-0000-0000-000000000000'
      );
      
      if (!result || !result.encounter) {
        throw new Error('Invalid advance response');
      }

      const { encounter, stats } = result;
      
      send({ 
        type: 'UPDATE_STATS', 
        o2: stats?.current_o2 ?? o2,
        fuel: stats?.current_fuel ?? fuel,
        encounters: [...(encounters || []), encounter],
        currentEncounter: encounter,
        expeditionId: currentExpeditionId,
        title: expeditionTitle
      });

      // Update current node in timeline
      if (currentIndex !== -1 && currentIndex < timeline.length - 1) {
        send({ type: 'NEXT_NODE' });
      } else if (currentIndex === timeline.length - 1) {
        send({ type: 'MISSION_END' });
      }

      if (stats && stats.current_o2 <= 0) {
        send({ type: 'MISSION_END' });
        setIsTransitioning(false);
        return;
      }

      if (encounter.type === 'COMBAT') {
        setTimeout(() => {
          send({ type: 'ENTER_COMBAT', enemyId: encounter.enemy_id || null });
          setIsTransitioning(false);
        }, 2000);
      } else {
        setIsTransitioning(false);
      }
    } catch (error) {
      console.error('Failed to advance timeline:', error);
      setIsTransitioning(false);
    }
  };

  const returnToBastion = () => {
    send({ type: 'RETURN' });
    if (user?.active_character_id) {
      bastionSystem.refreshVehicles(user.active_character_id);
    }
  };

  if (authLoading) {
    return (
      <div className="h-screen w-screen bg-black flex items-center justify-center text-white font-mono">
        <div className="flex flex-col items-center gap-4">
          <div className="w-12 h-12 border-4 border-pink-500 border-t-transparent rounded-full animate-spin" />
          <div className="text-xs uppercase tracking-[0.2em] animate-pulse">Synchronizing Neural Link...</div>
        </div>
      </div>
    );
  }

  return (
    <main className="relative h-screen w-screen overflow-hidden bg-black text-zinc-100 font-mono">
      <AnimatePresence mode="wait">
        {state.matches('landing') && (
          <motion.div key="landing" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <LandingStage onStart={() => send({ type: 'START', user })} />
          </motion.div>
        )}

        {state.matches('characterCreation') && (
          <motion.div key="char-creation" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <CharacterCreationStage onSuccess={() => send({ type: 'SUCCESS' })} />
          </motion.div>
        )}

        {state.matches('bastion') && (
          <motion.div key="bastion" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <Bastion 
              onDeploy={() => send({ type: 'DEPLOY' })} 
              onGacha={() => send({ type: 'OPEN_GACHA' })}
            />
          </motion.div>
        )}

        {state.matches('gacha') && (
          <motion.div key="gacha" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <GachaStage onBack={() => send({ type: 'BACK' })} />
          </motion.div>
        )}

        {state.matches('map') && (
          <motion.div key="map" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <UniverseMap 
              sectors={sectors}
              selectedSector={selectedSector}
              onSelectSector={(sector) => send({ type: 'SELECT_SECTOR', sector })}
              onScanSector={() => send({ type: 'SCAN' })}
              onBack={() => send({ type: 'BACK' })}
              selectedVehicleClass={selectedVehicle?.class}
            />
          </motion.div>
        )}

        {state.matches('locationScan') && selectedSector && (
          <motion.div key="location-scan" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <LocationScan 
              selectedSector={selectedSector}
              selectedSubSector={selectedSubSector}
              onSelectSubSector={(subSector) => send({ type: 'SELECT_SUBSECTOR', subSector })}
              onConfirmDeployment={confirmDeployment}
              onBack={() => send({ type: 'BACK' })}
              selectedVehicle={selectedVehicle}
              onSelectVehicle={(vehicle) => send({ type: 'SELECT_VEHICLE', vehicle })}
              upgrades={upgrades}
              inventory={inventory}
              canDeploy={canDeploy}
              meetsRequirements={meetsRequirements}
            />
          </motion.div>
        )}

        {state.matches('planetSurface') && selectedSubSector && (
          <motion.div key="planet-surface" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <PlanetSurface 
              selectedSubSector={selectedSubSector}
              selectedPlanetLocation={selectedPlanetLocation}
              onSelectPlanetLocation={(location) => send({ type: 'SELECT_PLANET_LOCATION', location })}
              onConfirmDeployment={confirmPlanetDeployment}
              onBack={() => send({ type: 'BACK' })}
              vehicles={vehicles}
              selectedVehicle={selectedVehicle}
              onSelectVehicle={(vehicle) => send({ type: 'SELECT_VEHICLE', vehicle })}
              inventory={inventory}
              canDeploy={canDeploy}
              meetsRequirements={meetsRequirements}
            />
          </motion.div>
        )}

        {state.matches('exploration') && (
          <motion.div key="exploration" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <TimelineView 
              expeditionTitle={expeditionTitle}
              o2={o2}
              fuel={fuel}
              timeline={timeline}
              currentNode={currentNode}
              onResolveChoice={resolveChoice}
              onEnterCombat={(enemyId) => {
                send({ type: 'ENTER_COMBAT', enemyId });
              }}
              onAdvance={advanceTimeline}
            />
          </motion.div>
        )}

        {state.matches('combat') && (
          <motion.div key="combat" initial={{ opacity: 0, scale: 0.9 }} animate={{ opacity: 1, scale: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <CombatStage 
              attackerId={selectedVehicle?.id || '00000000-0000-0000-0000-000000000000'}
              enemyId={activeEnemyId || ''}
              onCombatEnd={async (result) => {
                console.log('Combat ended:', result);
                if (result === 'VICTORY' && currentNode) {
                  try {
                    const updatedNode = await explorationSystem.resolveNode(currentNode.id);
                    send({ type: 'RESOLVE_NODE', node: updatedNode });
                  } catch (error) {
                    console.error('Failed to resolve node after combat:', error);
                  }
                }
                send({ type: 'COMBAT_END' });
              }}
            />
          </motion.div>
        )}

        {state.matches('debrief') && (
          <motion.div key="debrief" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full flex flex-col items-center justify-center p-8">
            <div className="max-w-2xl w-full border border-zinc-800 p-12 bg-zinc-900/10 backdrop-blur-md">
              <h2 className="text-4xl font-black italic mb-8 tracking-tighter">MISSION DEBRIEF</h2>
              <div className="space-y-4 mb-12">
                <div className="flex justify-between border-b border-zinc-900 pb-2">
                  <span className="text-zinc-500 uppercase text-xs">Distance Traveled</span>
                  <span className="font-bold">{encounters.length} Encounters</span>
                </div>
                <div className="flex justify-between border-b border-zinc-900 pb-2">
                  <span className="text-zinc-500 uppercase text-xs">Combat Encounters</span>
                  <span className="font-bold text-red-500">{encounters.filter(e => e.type === 'COMBAT').length}</span>
                </div>
              </div>
              <button 
                onClick={returnToBastion}
                className="w-full border border-white py-4 font-black uppercase tracking-tighter hover:bg-white hover:text-black transition-all"
              >
                Return to Mothership
              </button>
            </div>
          </motion.div>
        )}
      </AnimatePresence>

      <NotificationSystem />

      {/* SCANLINES EFFECT */}
      <div className="fixed inset-0 pointer-events-none opacity-[0.03] bg-[linear-gradient(rgba(18,16,16,0)_50%,rgba(0,0,0,0.25)_50%),linear-gradient(90deg,rgba(255,0,0,0.06),rgba(0,255,0,0.02),rgba(0,0,255,0.06))] z-50 bg-[length:100%_2px,3px_100%]" />
    </main>
  );
}
