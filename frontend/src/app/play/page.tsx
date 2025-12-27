'use client';

import { useState, useEffect } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { explorationService, Sector, SubSector, PlanetLocation } from '@/services/exploration';
import { useAuthSync } from '@/hooks/use-auth-sync';

// Components
import Bastion from '@/components/game/Bastion';
import UniverseMap from '@/components/game/UniverseMap';
import LocationScan from '@/components/game/LocationScan';
import PlanetSurface from '@/components/game/PlanetSurface';
import ExplorationLoop from '@/components/game/ExplorationLoop';
import CombatStage from '@/components/game/CombatStage';

type GameStage = 'BASTION' | 'MAP' | 'LOCATION_SCAN' | 'PLANET_SURFACE' | 'EXPLORATION' | 'COMBAT' | 'DEBRIEF';
type DeploymentMode = 'PILOT' | 'SPEEDER' | 'MECH' | 'TANK' | 'SHIP' | 'EXOSUIT' | 'HAULER';

interface BastionModules {
  radar: number;
  lab: number;
  warpDrive: number;
  atmosphericEntry: boolean;
  quantumGate: boolean;
}

export default function PlayPage() {
  const { user, isLoading: authLoading } = useAuthSync();
  const [stage, setStage] = useState<GameStage>('BASTION');

  // Game State
  const [o2, setO2] = useState(100);
  const [fuel, setFuel] = useState(100);
  const [scrapMetal, setScrapMetal] = useState(0);
  const [researchData, setResearchData] = useState(0);
  const [activeEnemyId, setActiveEnemyId] = useState<string | null>(null);
  const [encounters, setEncounters] = useState<any[]>([]);
  const [timelineNodes, setTimelineNodes] = useState<any[]>([]);
  const [currentEncounter, setCurrentEncounter] = useState<any | null>(null);
  const [expeditionTitle, setExpeditionTitle] = useState('THE SILENT SIGNAL');

  const [vehicles, setVehicles] = useState<any[]>([]);
  const [selectedVehicle, setSelectedVehicle] = useState<any | null>(null);
  const [sectors, setSectors] = useState<Sector[]>([]);
  const [selectedSector, setSelectedSector] = useState<Sector | null>(null);
  const [selectedSubSector, setSelectedSubSector] = useState<SubSector | null>(null);
  const [selectedPlanetLocation, setSelectedPlanetLocation] = useState<PlanetLocation | null>(null);
  const [currentExpeditionId, setCurrentExpeditionId] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);
  const [isTransitioning, setIsTransitioning] = useState(false);

  // Bastion & Inventory
  const [inventory, setInventory] = useState<string[]>(['Basic O2 Tank']);
  const [modules, setModules] = useState<BastionModules>({
    radar: 1,
    lab: 1,
    warpDrive: 1,
    atmosphericEntry: false,
    quantumGate: false
  });

  // Fetch Universe Map
  useEffect(() => {
    const fetchMap = async () => {
      try {
        const data = await explorationService.getUniverseMap();
        setSectors(data);
        setLoading(false);
      } catch (error) {
        console.error('Failed to fetch universe map:', error);
        setLoading(false);
      }
    };
    fetchMap();
  }, []);

  // Fetch Real Vehicles from DB
  useEffect(() => {
    const fetchVehicles = async () => {
      if (!user?.id) return;
      try {
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/vehicles?user_id=${user.id}`);
        const data = await response.json();
        if (Array.isArray(data)) {
          const mappedVehicles = data.map((m: any) => {
            const vehicleClass = m.class || m.model || 'UNKNOWN';
            const vehicleType = m.vehicle_type || (vehicleClass.includes('TANK') ? 'TANK' : vehicleClass.includes('SHIP') ? 'SHIP' : 'VEHICLE');

            return {
              id: m.id,
              class: vehicleClass,
              type: vehicleType as DeploymentMode,
              rarity: 'COMMON',
              stats: {
                hp: m.hp,
                attack: m.attack,
                defense: m.defense,
                speed: m.speed
              }
            };
          });
          setVehicles(mappedVehicles);
          if (mappedVehicles.length > 0) setSelectedVehicle(mappedVehicles[0]);
        }
      } catch (error) {
        console.error('Failed to fetch vehicles:', error);
      }
    };
    fetchVehicles();
  }, [user?.id]);

  // Temporary fix for missing upgrades state
  const upgrades = {
    atmosphericEntry: true,
    quantumGate: false,
    miningDrill: true,
    hackingModule: true
  };

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
    if (!selectedSubSector) return;

    if (!selectedVehicle) {
      alert("Please select a vehicle before deploying.");
      return;
    }

    if (selectedSubSector.type === 'PLANET' && selectedSubSector.locations) {
      setStage('PLANET_SURFACE');
      setSelectedPlanetLocation(null);
    } else {
      try {
        setIsTransitioning(true);
        if (!user?.id) throw new Error('User not authenticated');
        const result = await explorationService.startExploration(
          user.id,
          selectedSubSector.id,
          selectedVehicle?.id || '00000000-0000-0000-0000-000000000000'
        );

        setCurrentExpeditionId(result.expedition.id);
        setExpeditionTitle(`${selectedSector?.name} // ${selectedSubSector.name}`);

        // Fetch Timeline
        const timeline = await explorationService.getTimeline(result.expedition.id);
        setTimelineNodes(timeline);

        setStage('EXPLORATION');

        if (result.pilot_stats) {
          setO2(result.pilot_stats.current_o2);
          setFuel(result.pilot_stats.current_fuel);
          setScrapMetal(result.pilot_stats.scrap_metal);
          setResearchData(result.pilot_stats.research_data);
        }

        setEncounters(result.encounters);
        setCurrentEncounter(result.encounters[0]);
      } catch (error) {
        console.error('Failed to start exploration:', error);
      } finally {
        setIsTransitioning(false);
      }
    }
  };

  const startStoryMission = async (blueprintId: string) => {
    try {
      setIsTransitioning(true);
      if (!user?.id) throw new Error('User not authenticated');

      const result = await explorationService.startExploration(
        user.id,
        '00000000-0000-0000-0000-000000000000', // Dummy subsector for story
        selectedVehicle?.id || '00000000-0000-0000-0000-000000000000',
        undefined,
        blueprintId
      );

      setCurrentExpeditionId(result.expedition.id);
      setExpeditionTitle(result.expedition.title);
      setEncounters(result.encounters);
      setCurrentEncounter(result.encounters[0]);

      // Fetch Timeline
      const timeline = await explorationService.getTimeline(result.expedition.id);
      setTimelineNodes(timeline);

      setStage('EXPLORATION');

      if (result.pilot_stats) {
        setO2(result.pilot_stats.current_o2);
        setFuel(result.pilot_stats.current_fuel);
        setScrapMetal(result.pilot_stats.scrap_metal);
        setResearchData(result.pilot_stats.research_data);
      }

      setEncounters(result.encounters);
      setCurrentEncounter(result.encounters[0]);
    } catch (error) {
      console.error('Failed to start story mission:', error);
      alert('Failed to start story mission: ' + error);
    } finally {
      setIsTransitioning(false);
    }
  };

  const confirmPlanetDeployment = async () => {
    if (!selectedPlanetLocation || !selectedSubSector) return;

    if (!selectedVehicle) {
      alert("Please select a vehicle before deploying.");
      return;
    }

    try {
      setIsTransitioning(true);
      if (!user?.id) throw new Error('User not authenticated');
      const result = await explorationService.startExploration(
        user.id,
        selectedSubSector.id,
        selectedVehicle?.id || '00000000-0000-0000-0000-000000000000',
        selectedPlanetLocation.id
      );

      setCurrentExpeditionId(result.expedition.id);
      setExpeditionTitle(`${selectedSector?.name} // ${selectedSubSector?.name} // ${selectedPlanetLocation.name}`);

      // Fetch Timeline
      const timeline = await explorationService.getTimeline(result.expedition.id);
      setTimelineNodes(timeline);

      setStage('EXPLORATION');

      if (result.pilot_stats) {
        setO2(result.pilot_stats.current_o2);
        setFuel(result.pilot_stats.current_fuel);
      }

      setEncounters(result.encounters);
      setCurrentEncounter(result.encounters[0]);
    } catch (error) {
      console.error('Failed to start planet exploration:', error);
    } finally {
      setIsTransitioning(false);
    }
  };

  const advanceTimeline = async () => {
    if (isTransitioning || !currentExpeditionId) return;

    try {
      setIsTransitioning(true);
      const nextEncounterData = await explorationService.advanceTimeline(
        currentExpeditionId,
        selectedVehicle?.id || '00000000-0000-0000-0000-000000000000'
      );

      if (nextEncounterData.pilot_stats) {
        setO2(nextEncounterData.pilot_stats.current_o2);
        setFuel(nextEncounterData.pilot_stats.current_fuel);
        setScrapMetal(nextEncounterData.pilot_stats.scrap_metal);
        setResearchData(nextEncounterData.pilot_stats.research_data);

        if (nextEncounterData.pilot_stats.current_o2 <= 0) {
          setStage('DEBRIEF');
          return;
        }
      }

      const nextEncounter = nextEncounterData.encounter;

      // Refresh timeline to reflect resolved status in HUD
      const updatedTimeline = await explorationService.getTimeline(currentExpeditionId);
      setTimelineNodes(updatedTimeline);

      setEncounters(prev => [...prev, nextEncounter]);
      setCurrentEncounter(nextEncounter);

      if (nextEncounter.type === 'COMBAT') {
        setActiveEnemyId(nextEncounter.enemy_id || null);
        setTimeout(() => {
          setStage('COMBAT');
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

  return (
    <main className="relative h-screen w-screen overflow-hidden bg-black text-zinc-100 font-mono">
      <AnimatePresence mode="wait">
        {stage === 'BASTION' && (
          <motion.div key="bastion" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <Bastion
              onDeploy={() => setStage('MAP')}
              onGacha={() => { }}
              onResearch={() => { }}
              onStartStory={() => startStoryMission('iron-awakening')}
            />
          </motion.div>
        )}

        {stage === 'MAP' && (
          <motion.div key="map" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <UniverseMap
              sectors={sectors}
              selectedSector={selectedSector}
              onSelectSector={setSelectedSector}
              onScanSector={() => setStage('LOCATION_SCAN')}
              onBack={() => setStage('BASTION')}
              selectedVehicleClass={selectedVehicle?.class}
            />
          </motion.div>
        )}

        {stage === 'LOCATION_SCAN' && selectedSector && (
          <motion.div key="location-scan" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <LocationScan
              selectedSector={selectedSector}
              selectedSubSector={selectedSubSector}
              onSelectSubSector={setSelectedSubSector}
              onConfirmDeployment={confirmDeployment}
              onBack={() => setStage('MAP')}
              vehicles={vehicles}
              selectedVehicle={selectedVehicle}
              onSelectVehicle={setSelectedVehicle}
              upgrades={upgrades}
              inventory={inventory}
              canDeploy={canDeploy}
              meetsRequirements={meetsRequirements}
            />
          </motion.div>
        )}

        {stage === 'PLANET_SURFACE' && selectedSubSector && (
          <motion.div key="planet-surface" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <PlanetSurface
              selectedSubSector={selectedSubSector}
              selectedPlanetLocation={selectedPlanetLocation}
              onSelectPlanetLocation={setSelectedPlanetLocation}
              onConfirmDeployment={confirmPlanetDeployment}
              onBack={() => setStage('LOCATION_SCAN')}
              vehicles={vehicles}
              selectedVehicle={selectedVehicle}
              onSelectVehicle={setSelectedVehicle}
              inventory={inventory}
              canDeploy={canDeploy}
              meetsRequirements={meetsRequirements}
            />
          </motion.div>
        )}

        {stage === 'EXPLORATION' && (
          <motion.div key="exploration" initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <ExplorationLoop
              expeditionTitle={expeditionTitle}
              o2={o2}
              fuel={fuel}
              scrapMetal={scrapMetal}
              researchData={researchData}
              encounters={encounters}
              timelineNodes={timelineNodes}
              currentEncounter={currentEncounter}
              isTransitioning={isTransitioning}
              onAdvance={advanceTimeline}
              onEnterCombat={(enemyId) => {
                setActiveEnemyId(enemyId);
                setStage('COMBAT');
              }}
              onResolveNode={async (nodeId, choiceId) => {
                try {
                  if (choiceId) {
                    const result = await explorationService.resolveChoice(nodeId, choiceId);
                    if (result.pilot_stats) {
                      setO2(result.pilot_stats.current_o2);
                      setFuel(result.pilot_stats.current_fuel);
                      setScrapMetal(result.pilot_stats.scrap_metal);
                      setResearchData(result.pilot_stats.research_data);
                    }
                  } else {
                    await explorationService.resolveNode(nodeId);
                  }
                  // Refresh timeline
                  const updatedTimeline = await explorationService.getTimeline(currentExpeditionId!);
                  setTimelineNodes(updatedTimeline);
                } catch (error) {
                  console.error('Failed to resolve node:', error);
                  throw error;
                }
              }}
            />
          </motion.div>
        )}

        {stage === 'COMBAT' && (
          <motion.div key="combat" initial={{ opacity: 0, scale: 0.9 }} animate={{ opacity: 1, scale: 1 }} exit={{ opacity: 0 }} className="h-full w-full">
            <CombatStage
              attackerId={selectedVehicle?.id || '00000000-0000-0000-0000-000000000000'}
              enemyId={activeEnemyId || ''}
              onCombatEnd={async (result) => {
                console.log('Combat ended:', result);
                if (result === 'VICTORY') {
                  // Mark the current node as resolved
                  const currentNode = timelineNodes[encounters.length - 1];
                  if (currentNode) {
                    try {
                      await explorationService.resolveNode(currentNode.id);
                      // Refresh timeline
                      const updatedTimeline = await explorationService.getTimeline(currentExpeditionId!);
                      setTimelineNodes(updatedTimeline);
                    } catch (error) {
                      console.error('Failed to resolve node after combat:', error);
                    }
                  }
                }
                setStage('EXPLORATION');
              }}
            />
          </motion.div>
        )}

        {stage === 'DEBRIEF' && (
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
                <div className="flex justify-between border-b border-zinc-900 pb-2">
                  <span className="text-zinc-500 uppercase text-xs">Total Scrap Metal</span>
                  <span className="font-bold text-yellow-500">{scrapMetal}</span>
                </div>
                <div className="flex justify-between border-b border-zinc-900 pb-2">
                  <span className="text-zinc-500 uppercase text-xs">Total Research Data</span>
                  <span className="font-bold text-cyan-500">{researchData}</span>
                </div>
              </div>
              <button
                onClick={() => setStage('BASTION')}
                className="w-full border border-white py-4 font-black uppercase tracking-tighter hover:bg-white hover:text-black transition-all"
              >
                Return to Mothership
              </button>
            </div>
          </motion.div>
        )}
      </AnimatePresence>

      {/* SCANLINES EFFECT */}
      <div className="fixed inset-0 pointer-events-none opacity-[0.03] bg-[linear-gradient(rgba(18,16,16,0)_50%,rgba(0,0,0,0.25)_50%),linear-gradient(90deg,rgba(255,0,0,0.06),rgba(0,255,0,0.02),rgba(0,0,255,0.06))] z-50 bg-[length:100%_2px,3px_100%]" />
    </main>
  );
}
