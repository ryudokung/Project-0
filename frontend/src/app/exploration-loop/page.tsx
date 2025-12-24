'use client';

import { useState, useEffect } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { explorationService, Sector, SubSector, PlanetLocation } from '@/services/exploration';
import { useAuthSync } from '@/hooks/useAuthSync';

type GameState = 'HANGAR' | 'MAP' | 'LOCATION_SCAN' | 'PLANET_SURFACE' | 'EXPLORATION' | 'ENCOUNTER' | 'DEBRIEF';
type EncounterType = 'COMBAT' | 'RESOURCE' | 'NARRATIVE' | 'ANCHOR';
type POIType = 'PLANET' | 'STATION' | 'WRECK' | 'ANOMALY';
type DeploymentMode = 'PILOT' | 'SPEEDER' | 'MECH' | 'TANK' | 'SHIP' | 'EXOSUIT' | 'HAULER';

interface MothershipUpgrades {
  atmosphericEntry: boolean;
  quantumGate: boolean; // Renamed from teleport
  miningDrill: boolean;
  hackingModule: boolean;
  radarLevel: number;
  scannerLevel: number;
}

interface Encounter {
  id: string;
  type: EncounterType;
  title: string;
  description: string;
  visualPrompt: string;
  image: string;
}

interface Vehicle {
  id: string;
  class: string;
  type: DeploymentMode;
  rarity: string;
  stats: {
    hp: number;
    attack: number;
    defense: number;
    speed: number;
  };
  image_url?: string;
}

export default function ExplorationLoop() {
  const { user, isLoading: authLoading } = useAuthSync();
  const [gameState, setGameState] = useState<GameState>('HANGAR');
  const [o2, setO2] = useState(100);
  const [fuel, setFuel] = useState(100);
  const [encounters, setEncounters] = useState<Encounter[]>([]);
  const [currentEncounter, setCurrentEncounter] = useState<Encounter | null>(null);
  const [expeditionTitle, setExpeditionTitle] = useState('THE SILENT SIGNAL');
  
  const [vehicles, setVehicles] = useState<Vehicle[]>([]);
  const [selectedVehicle, setSelectedVehicle] = useState<Vehicle | null>(null);
  const [sectors, setSectors] = useState<Sector[]>([]);
  const [selectedSector, setSelectedSector] = useState<Sector | null>(null);
  const [selectedSubSector, setSelectedSubSector] = useState<SubSector | null>(null);
  const [selectedPlanetLocation, setSelectedPlanetLocation] = useState<PlanetLocation | null>(null);
  const [currentExpeditionId, setCurrentExpeditionId] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);
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

  const [currentMode, setCurrentMode] = useState<DeploymentMode>('PILOT');

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

  // Derived mode based on selection
  useEffect(() => {
    if (selectedVehicle) {
      setCurrentMode(selectedVehicle.type as DeploymentMode);
    } else {
      setCurrentMode('PILOT');
    }
  }, [selectedVehicle]);

  const canDeploy = (target: { allowedModes: string[], requiresAtmosphere: boolean }) => {
    // Check Mode
    const modeAllowed = target.allowedModes.includes(currentMode);
    
    // Check Entry System
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


  // Fetch Real Vehicles from DB
  useEffect(() => {
    const fetchVehicles = async () => {
      if (!user?.id) return;
      try {
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/mechs?user_id=${user.id}`);
        const data = await response.json();
        if (Array.isArray(data)) {
          // Map backend data to our Vehicle interface
          const mappedVehicles: Vehicle[] = data.map((m: any) => ({
            id: m.id,
            class: m.model,
            type: (m.model.includes('TANK') ? 'TANK' : m.model.includes('SHIP') ? 'SHIP' : 'MECH') as DeploymentMode,
            rarity: 'COMMON',
            stats: {
              hp: m.hp,
              attack: m.attack,
              defense: m.defense,
              speed: m.speed
            }
          }));
          setVehicles(mappedVehicles);
          if (mappedVehicles.length > 0) setSelectedVehicle(mappedVehicles[0]);
        }
      } catch (error) {
        console.error('Failed to fetch vehicles:', error);
        // Fallback mock vehicles for variety
        setVehicles([
          {
            id: 'mock-1',
            class: 'STRIKER',
            type: 'MECH',
            rarity: 'RARE',
            stats: { hp: 100, attack: 20, defense: 15, speed: 10 },
            image_url: 'https://images.unsplash.com/photo-1550684848-fac1c5b4e853?q=80&w=1000&auto=format&fit=crop'
          },
          {
            id: 'mock-2',
            class: 'GHOST',
            type: 'SPEEDER',
            rarity: 'EPIC',
            stats: { hp: 40, attack: 5, defense: 5, speed: 30 },
            image_url: 'https://images.unsplash.com/photo-1558981285-6f0c94958bb6?q=80&w=1000&auto=format&fit=crop'
          },
          {
            id: 'mock-3',
            class: 'IRON TUSK',
            type: 'TANK',
            rarity: 'RARE',
            stats: { hp: 200, attack: 35, defense: 40, speed: 5 },
            image_url: 'https://images.unsplash.com/photo-1599408162172-75232f09e9d1?q=80&w=1000&auto=format&fit=crop'
          },
          {
            id: 'mock-4',
            class: 'VOID RUNNER',
            type: 'SHIP',
            rarity: 'LEGENDARY',
            stats: { hp: 80, attack: 25, defense: 10, speed: 50 },
            image_url: 'https://images.unsplash.com/photo-1581822261290-991b38693d1b?q=80&w=1000&auto=format&fit=crop'
          },
          {
            id: 'mock-5',
            class: 'BULLDOG',
            type: 'HAULER',
            rarity: 'COMMON',
            stats: { hp: 150, attack: 10, defense: 20, speed: 8 },
            image_url: 'https://images.unsplash.com/photo-1519003722824-194d4455a60c?q=80&w=1000&auto=format&fit=crop'
          },
          {
            id: 'mock-6',
            class: 'SKELETON',
            type: 'EXOSUIT',
            rarity: 'EPIC',
            stats: { hp: 60, attack: 15, defense: 10, speed: 15 },
            image_url: 'https://images.unsplash.com/photo-1531259683007-016a7b628fc3?q=80&w=1000&auto=format&fit=crop'
          }
        ]);
      }
    };
    fetchVehicles();
  }, []);

  const startExploration = () => {
    setGameState('MAP');
  };

  const openLocationScan = () => {
    if (selectedSector) {
      setGameState('LOCATION_SCAN');
      setSelectedSubSector(null);
    }
  };

  const confirmDeployment = async () => {
    if (!selectedSubSector) return;
    
    if (selectedSubSector.type === 'PLANET' && selectedSubSector.locations) {
      setGameState('PLANET_SURFACE');
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
        setGameState('EXPLORATION');
        
        // Update stats from backend
        if (result.pilot_stats) {
          setO2(result.pilot_stats.current_o2);
          setFuel(result.pilot_stats.current_fuel);
        }

        const mappedEncounters = result.encounters.map((e: any) => ({
          id: e.id,
          type: e.type,
          title: e.title,
          description: e.description,
          visualPrompt: e.visual_prompt,
          image: 'https://images.unsplash.com/photo-1614728263952-84ea256f9679?q=80&w=1000&auto=format&fit=crop'
        }));
        
        setEncounters(mappedEncounters);
        setCurrentEncounter(mappedEncounters[0]);
      } catch (error) {
        console.error('Failed to start exploration:', error);
      } finally {
        setIsTransitioning(false);
      }
    }
  };

  const confirmPlanetDeployment = async () => {
    if (!selectedPlanetLocation || !selectedSubSector) return;
    
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
      setGameState('EXPLORATION');
      
      // Update stats from backend
      if (result.pilot_stats) {
        setO2(result.pilot_stats.current_o2);
        setFuel(result.pilot_stats.current_fuel);
      }

const mappedEncounters = result.encounters.map((e: any) => ({
          id: e.id,
          type: e.type,
          title: e.title,
          description: e.description,
          visualPrompt: e.visual_prompt,
        image: 'https://images.unsplash.com/photo-1614728263952-84ea256f9679?q=80&w=1000&auto=format&fit=crop'
      }));
      
        setEncounters(mappedEncounters);
        setCurrentEncounter(mappedEncounters[0]);
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
      
      // Call backend to advance
      const nextEncounterData = await explorationService.advanceTimeline(
        currentExpeditionId,
        selectedVehicle?.id || '00000000-0000-0000-0000-000000000000'
      );

      // Update stats from backend
      if (nextEncounterData.pilot_stats) {
        setO2(nextEncounterData.pilot_stats.current_o2);
        setFuel(nextEncounterData.pilot_stats.current_fuel);
        
        if (nextEncounterData.pilot_stats.current_o2 <= 0) {
          setGameState('DEBRIEF');
          return;
        }
      }

      const nextEncounter = {
        id: nextEncounterData.encounter.id,
        type: nextEncounterData.encounter.type,
        title: nextEncounterData.encounter.title,
        description: nextEncounterData.encounter.description,
        visualPrompt: nextEncounterData.encounter.visual_prompt,
        image: 'https://images.unsplash.com/photo-1550684848-fac1c5b4e853?q=80&w=1000&auto=format&fit=crop'
      };
      
      setEncounters(prev => [...prev, nextEncounter]);
      setCurrentEncounter(nextEncounter);

      if (nextEncounter.type === 'COMBAT') {
        // Show the combat encounter for a moment before switching to encounter
        setTimeout(() => {
          setGameState('ENCOUNTER');
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

  const finishEncounter = () => {
    setGameState('EXPLORATION');
    setIsTransitioning(false);
    // Update current encounter to reflect victory
    if (currentEncounter) {
      setCurrentEncounter({
        ...currentEncounter,
        type: 'NARRATIVE',
        title: 'Area Secured',
        description: 'The threat has been neutralized. The path forward is clear.'
      });
    }
  };

  return (
    <div className="min-h-screen bg-black text-white font-mono overflow-hidden selection:bg-pink-500 selection:text-white">
      
      {/* --- HANGAR STATE --- */}
      {gameState === 'HANGAR' && (
        <div className="flex flex-col items-center justify-center min-h-screen p-8">
          <motion.div 
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            className="border-2 border-zinc-800 p-8 md:p-12 bg-zinc-900/20 backdrop-blur-xl max-w-4xl w-full"
          >
            <div className="text-center mb-8">
              <h1 className="text-4xl font-black tracking-tighter mb-2 text-pink-500 italic">PROJECT-0</h1>
              <p className="text-zinc-500 uppercase tracking-widest text-xs">Mothership Hangar / Database Connected</p>
            </div>
            
            <div className="grid md:grid-cols-2 gap-8 mb-12">
              {/* Asset Inventory */}
              <div className="space-y-4">
                <h3 className="text-[10px] text-zinc-500 uppercase tracking-widest">Asset Inventory</h3>
                {loading ? (
                  <div className="text-zinc-600 animate-pulse">Accessing Database...</div>
                ) : (
                  <div className="space-y-2 max-h-64 overflow-y-auto pr-2">
                    {/* Pilot Status */}
                    <div className="p-4 border border-zinc-800 bg-zinc-900/20">
                      <div className="flex justify-between items-center">
                        <span className="font-bold italic text-pink-500">OPERATIVE STATUS</span>
                        <span className="text-[10px] bg-zinc-800 px-2 py-0.5 text-zinc-400">ACTIVE</span>
                      </div>
                      <div className="text-[10px] text-zinc-500 mt-1">Pilot EVA Gear: Standard Issue</div>
                    </div>

                    {vehicles.map((v, idx) => (
                      <div 
                        key={`vehicle-select-${v.id || idx}`}
                        onClick={() => setSelectedVehicle(v)}
                        className={`p-4 border cursor-pointer transition-all ${
                          selectedVehicle?.id === v.id ? 'border-pink-500 bg-pink-500/10' : 'border-zinc-800 hover:border-zinc-600'
                        }`}
                      >
                        <div className="flex justify-between items-center">
                          <span className="font-bold italic">{v.class}</span>
                          <span className="text-[10px] bg-zinc-800 px-2 py-0.5">{v.type}</span>
                        </div>
                        <div className="text-[10px] text-zinc-500 mt-1">{v.rarity} // ID: {v.id.slice(0,8)}...</div>
                      </div>
                    ))}
                  </div>
                )}
              </div>

              {/* Asset Details */}
              <div className="border border-zinc-800 p-6 bg-black/40">
                <h3 className="text-[10px] text-zinc-500 uppercase tracking-widest mb-4">Deployment Profile</h3>
                {selectedVehicle ? (
                  <div className="space-y-4">
                    <div className="aspect-video bg-zinc-900 border border-zinc-800 overflow-hidden">
                      {selectedVehicle.image_url ? (
                        <img src={selectedVehicle.image_url} className="w-full h-full object-cover" />
                      ) : (
                        <div className="w-full h-full flex items-center justify-center text-zinc-800 text-[10px]">NO VISUAL DNA</div>
                      )}
                    </div>
                    <div>
                      <div className="text-2xl font-black italic tracking-tighter">{selectedVehicle.class}</div>
                      <div className="text-[10px] text-pink-500 font-bold uppercase">{selectedVehicle.type} // {selectedVehicle.rarity}</div>
                    </div>
                    <div className="grid grid-cols-2 gap-2 text-[10px] border-t border-zinc-800 pt-4">
                      <div><span className="text-zinc-500">HP:</span> {selectedVehicle.stats.hp}</div>
                      <div><span className="text-zinc-500">ATK:</span> {selectedVehicle.stats.attack}</div>
                      <div><span className="text-zinc-500">DEF:</span> {selectedVehicle.stats.defense}</div>
                      <div><span className="text-zinc-500">SPD:</span> {selectedVehicle.stats.speed}</div>
                    </div>
                    <div className="text-[8px] text-zinc-600 italic mt-4">
                      * This asset is currently set as your primary deployment choice. You can switch assets during mission briefing.
                    </div>
                  </div>
                ) : (
                  <div className="h-full flex flex-col items-center justify-center text-zinc-700 italic text-sm text-center space-y-4">
                    <div className="w-12 h-12 border border-zinc-800 rounded-full flex items-center justify-center animate-pulse">
                      <div className="w-6 h-6 border border-zinc-700 rounded-full" />
                    </div>
                    <div>Select an asset to view technical specifications</div>
                  </div>
                )}
              </div>
            </div>

            <div className="mb-8 border border-zinc-800 p-6">
              <h2 className="text-xl font-bold mb-4 uppercase tracking-tighter">Mothership Research Tree</h2>
              <div className="grid grid-cols-2 gap-4">
                <div 
                  onClick={() => setUpgrades(prev => ({ ...prev, atmosphericEntry: !prev.atmosphericEntry }))}
                  className={`p-3 border cursor-pointer transition-all ${upgrades.atmosphericEntry ? 'border-green-500 bg-green-500/10' : 'border-zinc-800 text-zinc-600'}`}
                >
                  <div className="text-[10px] uppercase font-bold">Atmospheric Entry</div>
                  <div className="text-[8px]">{upgrades.atmosphericEntry ? 'ONLINE' : 'OFFLINE'}</div>
                </div>
                <div 
                  onClick={() => setUpgrades(prev => ({ ...prev, quantumGate: !prev.quantumGate }))}
                  className={`p-3 border cursor-pointer transition-all ${upgrades.quantumGate ? 'border-blue-500 bg-blue-500/10' : 'border-zinc-800 text-zinc-600'}`}
                >
                  <div className="text-[10px] uppercase font-bold">Quantum Gate</div>
                  <div className="text-[8px]">{upgrades.quantumGate ? 'ONLINE' : 'OFFLINE'}</div>
                </div>
                <div 
                  onClick={() => setUpgrades(prev => ({ ...prev, miningDrill: !prev.miningDrill }))}
                  className={`p-3 border cursor-pointer transition-all ${upgrades.miningDrill ? 'border-yellow-500 bg-yellow-500/10' : 'border-zinc-800 text-zinc-600'}`}
                >
                  <div className="text-[10px] uppercase font-bold">Mining Drill</div>
                  <div className="text-[8px]">{upgrades.miningDrill ? 'INSTALLED' : 'NOT INSTALLED'}</div>
                </div>
                <div 
                  onClick={() => setUpgrades(prev => ({ ...prev, hackingModule: !prev.hackingModule }))}
                  className={`p-3 border cursor-pointer transition-all ${upgrades.hackingModule ? 'border-purple-500 bg-purple-500/10' : 'border-zinc-800 text-zinc-600'}`}
                >
                  <div className="text-[10px] uppercase font-bold">Hacking Module</div>
                  <div className="text-[8px]">{upgrades.hackingModule ? 'INSTALLED' : 'NOT INSTALLED'}</div>
                </div>
              </div>
              {(upgrades.quantumGate || upgrades.miningDrill || upgrades.hackingModule) && (
                <div className="mt-4 space-y-1">
                  {upgrades.quantumGate && (
                    <div className="text-[8px] text-blue-400 italic uppercase">
                      Quantum Gate: Bypasses all atmospheric entry requirements.
                    </div>
                  )}
                  {upgrades.miningDrill && (
                    <div className="text-[8px] text-yellow-400 italic uppercase">
                      Mining Drill: Enables extraction at surface mining sites.
                    </div>
                  )}
                  {upgrades.hackingModule && (
                    <div className="text-[8px] text-purple-400 italic uppercase">
                      Hacking Module: Enables access to Syndicate bunkers.
                    </div>
                  )}
                </div>
              )}
            </div>

            <button 
              onClick={startExploration}
              className="w-full bg-white text-black py-4 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all duration-300"
            >
              Open Universe Map
            </button>
          </motion.div>
        </div>
      )}

      {/* --- MAP STATE --- */}
      {gameState === 'MAP' && (
        <div className="h-screen flex flex-col p-8">
          <div className="flex justify-between items-center mb-8">
            <div>
              <h2 className="text-3xl font-black italic tracking-tighter text-pink-500">UNIVERSE MAP</h2>
              <p className="text-zinc-500 text-xs uppercase tracking-[0.3em]">Select Destination Sector</p>
            </div>
            <div className="text-right">
              <div className="text-[10px] text-zinc-500 uppercase mb-1">Selected Vehicle</div>
              <div className="text-sm font-bold italic">{selectedVehicle?.class || 'PILOT'}</div>
            </div>
          </div>

          <div className="flex-1 grid md:grid-cols-3 gap-8">
            {/* Map Visualization */}
            <div className="md:col-span-2 border border-zinc-800 bg-zinc-900/10 relative overflow-hidden group">
              {/* Grid Background */}
              <div className="absolute inset-0 opacity-20" 
                style={{ backgroundImage: 'radial-gradient(circle, #333 1px, transparent 1px)', backgroundSize: '40px 40px' }} 
              />
              
              {/* Sector Points */}
              {sectors.map((s, idx) => (
                <motion.div
                  key={`sector-${s.id || idx}`}
                  initial={{ scale: 0 }}
                  animate={{ scale: 1 }}
                  whileHover={{ scale: 1.2 }}
                  onClick={() => setSelectedSector(s)}
                  className="absolute cursor-pointer group/point"
                  style={{ left: `${s.coordinates.x}%`, top: `${s.coordinates.y}%` }}
                >
                  <div className={`w-4 h-4 rounded-full border-2 border-white ${
                    selectedSector?.id === s.id ? 'bg-white scale-150' : 'bg-transparent'
                  } transition-all duration-300 shadow-[0_0_10px_rgba(255,255,255,0.5)]`} 
                  style={{ borderColor: s.color === 'pink' ? '#ec4899' : s.color === 'blue' ? '#3b82f6' : '#ef4444' }}
                  />
                  <div className={`absolute top-6 left-1/2 -translate-x-1/2 whitespace-nowrap text-[10px] font-black tracking-widest uppercase transition-opacity ${
                    selectedSector?.id === s.id ? 'opacity-100' : 'opacity-40 group-hover/point:opacity-100'
                  }`}
                  style={{ color: s.color === 'pink' ? '#ec4899' : s.color === 'blue' ? '#3b82f6' : '#ef4444' }}
                  >
                    {s.name}
                  </div>
                </motion.div>
              ))}

              {/* Scanning Line */}
              <motion.div 
                animate={{ top: ['0%', '100%', '0%'] }}
                transition={{ duration: 10, repeat: Infinity, ease: "linear" }}
                className="absolute left-0 w-full h-[1px] bg-pink-500/30 shadow-[0_0_15px_rgba(236,72,153,0.5)] z-0"
              />
            </div>

            {/* Sector Details */}
            <div className="flex flex-col gap-6">
              <div className="flex-1 border border-zinc-800 p-6 bg-black/40 backdrop-blur-md">
                <h3 className="text-[10px] text-zinc-500 uppercase tracking-widest mb-6">Sector Intelligence</h3>
                
                {selectedSector ? (
                  <motion.div 
                    initial={{ opacity: 0, x: 20 }}
                    animate={{ opacity: 1, x: 0 }}
                    className="space-y-6"
                  >
                    <div>
                      <div className="text-2xl font-black italic mb-2">{selectedSector.name}</div>
                      <p className="text-sm text-zinc-400 leading-relaxed">{selectedSector.description}</p>
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                      <div className="border border-zinc-800 p-3">
                        <div className="text-[8px] text-zinc-500 uppercase mb-1">Threat Level</div>
                        <div className={`text-sm font-bold ${
                          selectedSector.difficulty === 'HIGH' ? 'text-red-500' : 
                          selectedSector.difficulty === 'MEDIUM' ? 'text-pink-500' : 'text-blue-500'
                        }`}>{selectedSector.difficulty}</div>
                      </div>
                      <div className="border border-zinc-800 p-3">
                        <div className="text-[8px] text-zinc-500 uppercase mb-1">Coordinates</div>
                        <div className="text-sm font-mono">{selectedSector.coordinates.x}, {selectedSector.coordinates.y}</div>
                      </div>
                    </div>

                    <div className="pt-4">
                      <div className="text-[8px] text-zinc-500 uppercase mb-2">Visual DNA Signature</div>
                      <div className="h-24 bg-zinc-900 border border-zinc-800 flex items-center justify-center italic text-zinc-700 text-xs">
                        [ANALYZING ATMOSPHERIC DATA...]
                      </div>
                    </div>
                  </motion.div>
                ) : (
                  <div className="h-full flex items-center justify-center text-zinc-700 italic text-sm text-center">
                    Select a sector on the map to view intelligence data
                  </div>
                )}
              </div>

              <button 
                onClick={openLocationScan}
                disabled={!selectedSector}
                className="w-full bg-white text-black py-4 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all disabled:opacity-20"
              >
                Scan Sector
              </button>
              <button 
                onClick={() => setGameState('HANGAR')}
                className="w-full border border-zinc-800 text-zinc-500 py-2 text-xs uppercase tracking-widest hover:text-white transition-all"
              >
                Return to Hangar
              </button>
            </div>
          </div>
        </div>
      )}

      {/* --- LOCATION SCAN STATE --- */}
      {gameState === 'LOCATION_SCAN' && selectedSector && (
        <div className="h-screen flex flex-col p-8">
          <div className="flex justify-between items-center mb-8">
            <div>
              <h2 className="text-3xl font-black italic tracking-tighter text-pink-500">{selectedSector.name}</h2>
              <p className="text-zinc-500 text-xs uppercase tracking-[0.3em]">Sector Scan / Identify Points of Interest</p>
            </div>
            <button 
              onClick={() => setGameState('MAP')}
              className="text-zinc-500 hover:text-white text-xs uppercase tracking-widest"
            >
              Back to Universe Map
            </button>
          </div>

          <div className="flex-1 grid md:grid-cols-3 gap-8">
            {/* Local Map Visualization */}
            <div className="md:col-span-2 border border-zinc-800 bg-zinc-900/10 relative overflow-hidden">
              {/* Sector Visual Placeholder - Abstract Grid/Radar */}
              <div className="absolute inset-0 flex items-center justify-center">
                <div className="w-full h-full opacity-20" 
                  style={{ backgroundImage: 'linear-gradient(#333 1px, transparent 1px), linear-gradient(90deg, #333 1px, transparent 1px)', backgroundSize: '100px 100px' }} 
                />
                <motion.div 
                  animate={{ scale: [1, 1.1, 1], opacity: [0.1, 0.2, 0.1] }}
                  transition={{ duration: 4, repeat: Infinity }}
                  className="absolute w-[500px] h-[500px] rounded-full border border-zinc-800"
                />
              </div>
              
              {/* Sub-Sector Points */}
              {selectedSector.subSectors.map((ss, idx) => (
                <motion.div
                  key={`subsector-${ss.id || idx}`}
                  initial={{ scale: 0 }}
                  animate={{ scale: 1 }}
                  whileHover={{ scale: 1.2 }}
                  onClick={() => setSelectedSubSector(ss)}
                  className="absolute cursor-pointer group/point"
                  style={{ left: `${ss.coordinates.x}%`, top: `${ss.coordinates.y}%` }}
                >
                  {ss.type === 'PLANET' ? (
                    <div className={`w-12 h-12 rounded-full border-2 flex items-center justify-center ${
                      selectedSubSector?.id === ss.id ? 'bg-white text-black' : 'bg-transparent text-white'
                    } transition-all duration-300 shadow-[0_0_20px_rgba(255,255,255,0.2)]`}
                    style={{ borderColor: selectedSector.color === 'pink' ? '#ec4899' : selectedSector.color === 'blue' ? '#3b82f6' : '#ef4444' }}
                    >
                      <div className="w-8 h-8 rounded-full border border-white/20 animate-pulse" />
                    </div>
                  ) : (
                    <div className={`w-8 h-8 border flex items-center justify-center rotate-45 ${
                      selectedSubSector?.id === ss.id ? 'bg-white text-black' : 'bg-transparent text-white'
                    } transition-all duration-300`}
                    style={{ borderColor: selectedSector.color === 'pink' ? '#ec4899' : selectedSector.color === 'blue' ? '#3b82f6' : '#ef4444' }}
                    >
                      <div className="-rotate-45 text-[8px] font-black">
                        {ss.type === 'STATION' ? 'STA' : ss.type === 'WRECK' ? 'WRK' : 'POI'}
                      </div>
                    </div>
                  )}
                  <div className="absolute top-14 left-1/2 -translate-x-1/2 whitespace-nowrap text-[10px] font-black tracking-widest uppercase">
                    {ss.name}
                  </div>
                </motion.div>
              ))}
            </div>

            {/* Sub-Sector Details */}
            <div className="flex flex-col gap-6">
              <div className="flex-1 border border-zinc-800 p-6 bg-black/40 backdrop-blur-md overflow-y-auto">
                <h3 className="text-[10px] text-zinc-500 uppercase tracking-widest mb-6">Point of Interest Intel</h3>
                
                {selectedSubSector ? (
                  <motion.div 
                    initial={{ opacity: 0, x: 20 }}
                    animate={{ opacity: 1, x: 0 }}
                    className="space-y-6"
                  >
                    <div>
                      <div className="text-2xl font-black italic mb-2">{selectedSubSector.name}</div>
                      <p className="text-sm text-zinc-400 leading-relaxed">{selectedSubSector.description}</p>
                    </div>

                    {/* Suitability Analysis */}
                    <div className="space-y-3">
                      <div className="text-[8px] text-zinc-500 uppercase">Suitability Analysis</div>
                      <div className="space-y-2">
                        <div className="flex justify-between text-[10px]">
                          <span>Pilot Only</span>
                          <span className={selectedSubSector.suitability.pilot > 50 ? 'text-green-500' : 'text-red-500'}>
                            {selectedSubSector.suitability.pilot}%
                          </span>
                        </div>
                        <div className="w-full h-1 bg-zinc-900">
                          <div className="h-full bg-white" style={{ width: `${selectedSubSector.suitability.pilot}%` }} />
                        </div>
                        <div className="flex justify-between text-[10px]">
                          <span>Heavy Asset Deployment</span>
                          <span className={selectedSubSector.suitability.mech > 50 ? 'text-green-500' : 'text-red-500'}>
                            {selectedSubSector.suitability.mech}%
                          </span>
                        </div>
                        <div className="w-full h-1 bg-zinc-900">
                          <div className="h-full bg-pink-500" style={{ width: `${selectedSubSector.suitability.mech}%` }} />
                        </div>
                      </div>
                    </div>

                    {/* Asset Selection */}
                    <div className="space-y-3">
                      <div className="text-[8px] text-zinc-500 uppercase">Select Deployment Asset</div>
                      <div className="grid grid-cols-1 gap-2">
                        <div 
                          onClick={() => setSelectedVehicle(null)}
                          className={`p-2 border text-[10px] cursor-pointer transition-all ${!selectedVehicle ? 'border-white bg-white/10' : 'border-zinc-800 hover:border-zinc-700'}`}
                        >
                          <div className="flex justify-between">
                            <span className="font-bold">PILOT ONLY</span>
                            <span className={selectedSubSector.allowedModes.includes('PILOT') ? 'text-green-500' : 'text-red-500'}>
                              {selectedSubSector.allowedModes.includes('PILOT') ? '✓' : '✗'}
                            </span>
                          </div>
                        </div>
                        {vehicles.map((v, idx) => (
                          <div 
                            key={`vehicle-deploy-subsector-${v.id || idx}`}
                            onClick={() => setSelectedVehicle(v)}
                            className={`p-2 border text-[10px] cursor-pointer transition-all ${selectedVehicle?.id === v.id ? 'border-pink-500 bg-pink-500/10' : 'border-zinc-800 hover:border-zinc-700'}`}
                          >
                            <div className="flex justify-between">
                              <span className="font-bold">{v.class} ({v.type})</span>
                              <span className={selectedSubSector.allowedModes.includes(v.type) ? 'text-green-500' : 'text-red-500'}>
                                {selectedSubSector.allowedModes.includes(v.type) ? '✓' : '✗'}
                              </span>
                            </div>
                          </div>
                        ))}
                      </div>
                    </div>

                    {/* Deployment Compatibility */}
                    <div className="space-y-4">
                      <div className="text-[8px] text-zinc-500 uppercase">System Check</div>
                      
                      <div className="flex justify-between items-center border-b border-zinc-900 pb-2">
                        <span className="text-[10px] text-zinc-400">Entry System</span>
                        <span className={(!selectedSubSector.requiresAtmosphere || upgrades.atmosphericEntry || upgrades.quantumGate) ? 'text-green-500' : 'text-red-500'}>
                          {upgrades.quantumGate ? '✓ QUANTUM GATE READY' : 
                           (!selectedSubSector.requiresAtmosphere ? '✓ NO ENTRY REQ' : 
                           (upgrades.atmosphericEntry ? '✓ ATMOSPHERE READY' : '✗ ENTRY SYSTEM REQUIRED'))}
                        </span>
                      </div>
                    </div>

                    {/* Requirement Warning */}
                    {!canDeploy(selectedSubSector) && (
                      <div className="p-3 border border-red-500 bg-red-500/10 text-[10px] text-red-500 font-bold uppercase">
                        Deployment Blocked: Check Compatibility
                      </div>
                    )}

                    {/* Suitability Warning */}
                    {!selectedVehicle && selectedSubSector.suitability.pilot < 30 && (
                      <div className="p-3 border border-red-500/50 bg-red-500/10 text-[10px] text-red-500 font-bold animate-pulse">
                        WARNING: EXTREME RISK FOR PILOT-ONLY DEPLOYMENT
                      </div>
                    )}
                  </motion.div>
                ) : (
                  <div className="h-full flex items-center justify-center text-zinc-700 italic text-sm text-center">
                    Select a Landing Zone to analyze
                  </div>
                )}
              </div>

              <button 
                onClick={confirmDeployment}
                disabled={!selectedSubSector || !meetsRequirements(selectedSubSector.requirements) || !canDeploy(selectedSubSector)}
                className="w-full bg-white text-black py-4 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all disabled:opacity-20"
              >
                {selectedSubSector?.type === 'PLANET' ? 'Scan Surface' : 'Initiate Deployment'}
              </button>
            </div>
          </div>
        </div>
      )}

      {/* --- PLANET SURFACE STATE --- */}
      {gameState === 'PLANET_SURFACE' && selectedSubSector && (
        <div className="h-screen flex flex-col p-8">
          <div className="flex justify-between items-center mb-8">
            <div>
              <h2 className="text-3xl font-black italic tracking-tighter text-pink-500">{selectedSubSector.name}</h2>
              <p className="text-zinc-500 text-xs uppercase tracking-[0.3em]">Surface Scan / Select Tactical Objective</p>
            </div>
            <button 
              onClick={() => setGameState('LOCATION_SCAN')}
              className="text-zinc-500 hover:text-white text-xs uppercase tracking-widest"
            >
              Back to Sector Scan
            </button>
          </div>

          <div className="flex-1 grid md:grid-cols-3 gap-8">
            {/* Surface Map Visualization */}
            <div className="md:col-span-2 border border-zinc-800 bg-zinc-900/10 relative overflow-hidden">
              {/* Surface Visual Placeholder - Topographic Grid */}
              <div className="absolute inset-0 opacity-30" 
                style={{ 
                  backgroundImage: 'radial-gradient(circle at 50% 50%, #222 0%, transparent 70%)',
                  backgroundSize: '100% 100%'
                }} 
              />
              <div className="absolute inset-0 flex items-center justify-center">
                <div className="w-full h-full opacity-10" 
                  style={{ backgroundImage: 'linear-gradient(#444 1px, transparent 1px), linear-gradient(90deg, #444 1px, transparent 1px)', backgroundSize: '50px 50px' }} 
                />
              </div>
              
              {/* Planet Locations */}
              {selectedSubSector.locations?.map((loc, idx) => (
                <motion.div
                  key={`location-${loc.id || idx}`}
                  initial={{ scale: 0 }}
                  animate={{ scale: 1 }}
                  whileHover={{ scale: 1.2 }}
                  onClick={() => setSelectedPlanetLocation(loc)}
                  className="absolute cursor-pointer group/point"
                  style={{ left: `${loc.coordinates.x}%`, top: `${loc.coordinates.y}%` }}
                >
                  <div className={`w-10 h-10 border-2 flex items-center justify-center ${
                    selectedPlanetLocation?.id === loc.id ? 'bg-pink-500 text-white' : 'bg-transparent text-pink-500'
                  } transition-all duration-300 border-pink-500`}
                  >
                    <div className="text-[8px] font-black">OBJ</div>
                  </div>
                  <div className="absolute top-12 left-1/2 -translate-x-1/2 whitespace-nowrap text-[10px] font-black tracking-widest uppercase">
                    {loc.name}
                  </div>
                </motion.div>
              ))}
            </div>

            {/* Location Details */}
            <div className="flex flex-col gap-6">
              <div className="flex-1 border border-zinc-800 p-6 bg-black/40 backdrop-blur-md overflow-y-auto">
                <h3 className="text-[10px] text-zinc-500 uppercase tracking-widest mb-6">Objective Intelligence</h3>
                
                {selectedPlanetLocation ? (
                  <motion.div 
                    initial={{ opacity: 0, x: 20 }}
                    animate={{ opacity: 1, x: 0 }}
                    className="space-y-6"
                  >
                    <div>
                      <div className="text-2xl font-black italic mb-2">{selectedPlanetLocation.name}</div>
                      <p className="text-sm text-zinc-400 leading-relaxed">{selectedPlanetLocation.description}</p>
                    </div>

                    {/* Suitability Analysis */}
                    <div className="space-y-3">
                      <div className="text-[8px] text-zinc-500 uppercase">Tactical Suitability</div>
                      <div className="space-y-2">
                        <div className="flex justify-between text-[10px]">
                          <span>Pilot Stealth</span>
                          <span className={selectedPlanetLocation.suitability.pilot > 50 ? 'text-green-500' : 'text-red-500'}>
                            {selectedPlanetLocation.suitability.pilot}%
                          </span>
                        </div>
                        <div className="w-full h-1 bg-zinc-900">
                          <div className="h-full bg-white" style={{ width: `${selectedPlanetLocation.suitability.pilot}%` }} />
                        </div>
                        <div className="flex justify-between text-[10px]">
                          <span>Heavy Asset Power</span>
                          <span className={selectedPlanetLocation.suitability.mech > 50 ? 'text-green-500' : 'text-red-500'}>
                            {selectedPlanetLocation.suitability.mech}%
                          </span>
                        </div>
                        <div className="w-full h-1 bg-zinc-900">
                          <div className="h-full bg-pink-500" style={{ width: `${selectedPlanetLocation.suitability.mech}%` }} />
                        </div>
                      </div>
                    </div>

                    {/* Rewards & Requirements */}
                    <div className="grid grid-cols-2 gap-4">
                      <div>
                        <div className="text-[8px] text-zinc-500 uppercase mb-2">Objective Rewards</div>
                        <ul className="text-[10px] space-y-1">
                          {selectedPlanetLocation.rewards.map((r, idx) => <li key={`reward-${idx}-${r}`} className="text-blue-400">+ {r}</li>)}
                        </ul>
                      </div>
                      <div>
                        <div className="text-[8px] text-zinc-500 uppercase mb-2">Requirements</div>
                        <ul className="text-[10px] space-y-1">
                          {selectedPlanetLocation.requirements.map((r, idx) => (
                            <li key={`req-${idx}-${r}`} className={inventory.includes(r) ? 'text-green-500' : 'text-red-500 animate-pulse'}>
                              {inventory.includes(r) ? '✓' : '✗'} {r}
                            </li>
                          ))}
                        </ul>
                      </div>
                    </div>

                    {/* Asset Selection */}
                    <div className="space-y-3">
                      <div className="text-[8px] text-zinc-500 uppercase">Select Deployment Asset</div>
                      <div className="grid grid-cols-1 gap-2">
                        <div 
                          onClick={() => setSelectedVehicle(null)}
                          className={`p-2 border text-[10px] cursor-pointer transition-all ${!selectedVehicle ? 'border-white bg-white/10' : 'border-zinc-800 hover:border-zinc-700'}`}
                        >
                          <div className="flex justify-between">
                            <span className="font-bold">PILOT ONLY</span>
                            <span className={selectedPlanetLocation.allowedModes.includes('PILOT') ? 'text-green-500' : 'text-red-500'}>
                              {selectedPlanetLocation.allowedModes.includes('PILOT') ? '✓' : '✗'}
                            </span>
                          </div>
                        </div>
                        {vehicles.map((v, idx) => (
                          <div 
                            key={`vehicle-deploy-location-${v.id || idx}`}
                            onClick={() => setSelectedVehicle(v)}
                            className={`p-2 border text-[10px] cursor-pointer transition-all ${selectedVehicle?.id === v.id ? 'border-pink-500 bg-pink-500/10' : 'border-zinc-800 hover:border-zinc-700'}`}
                          >
                            <div className="flex justify-between">
                              <span className="font-bold">{v.class} ({v.type})</span>
                              <span className={selectedPlanetLocation.allowedModes.includes(v.type) ? 'text-green-500' : 'text-red-500'}>
                                {selectedPlanetLocation.allowedModes.includes(v.type) ? '✓' : '✗'}
                              </span>
                            </div>
                          </div>
                        ))}
                      </div>
                    </div>

                    {/* Requirement Warning */}
                    {(!meetsRequirements(selectedPlanetLocation.requirements) || !canDeploy(selectedPlanetLocation)) && (
                      <div className="p-3 border border-red-500 bg-red-500/10 text-[10px] text-red-500 font-bold uppercase">
                        Landing Blocked: Check Requirements
                      </div>
                    )}
                  </motion.div>
                ) : (
                  <div className="h-full flex items-center justify-center text-zinc-700 italic text-sm text-center">
                    Select a tactical objective on the surface
                  </div>
                )}
              </div>

              <button 
                onClick={confirmPlanetDeployment}
                disabled={!selectedPlanetLocation || !meetsRequirements(selectedPlanetLocation.requirements) || !canDeploy(selectedPlanetLocation)}
                className="w-full bg-white text-black py-4 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all disabled:opacity-20"
              >
                Initiate Deployment
              </button>
            </div>
          </div>
        </div>
      )}

      {/* --- EXPLORATION STATE --- */}
      {gameState === 'EXPLORATION' && (
        <div className="h-screen flex flex-col">
          {/* HUD TOP */}
          <div className="p-6 flex justify-between items-start border-b border-zinc-900 bg-black/80 backdrop-blur-md z-10">
            <div>
              <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Current Expedition</div>
              <div className="text-lg font-black italic text-pink-500">{expeditionTitle}</div>
            </div>
            <div className="flex gap-8">
              <div className="text-right">
                <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Oxygen Level</div>
                <div className={`text-xl font-black ${o2 < 30 ? 'text-red-500 animate-pulse' : 'text-white'}`}>{o2}%</div>
                <div className="w-32 h-1 bg-zinc-900 mt-1">
                  <motion.div className="h-full bg-white" animate={{ width: `${o2}%` }} />
                </div>
              </div>
              <div className="text-right">
                <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-1">Fuel Reserve</div>
                <div className="text-xl font-black">{fuel}%</div>
                <div className="w-32 h-1 bg-zinc-900 mt-1">
                  <motion.div className="h-full bg-zinc-500" animate={{ width: `${fuel}%` }} />
                </div>
              </div>
            </div>
          </div>

          {/* MAIN VIEWPORT */}
          <div 
            className="flex-1 relative flex items-center justify-center p-4 md:p-12 cursor-pointer group"
            onClick={() => {
              if (gameState === 'EXPLORATION') {
                if (currentEncounter?.type === 'COMBAT') {
                  // Manual skip to combat
                  setGameState('ENCOUNTER');
                  setIsTransitioning(false);
                } else {
                  advanceTimeline();
                }
              }
            }}
          >
            <AnimatePresence mode="wait">
              {currentEncounter && (
                <motion.div 
                  key={currentEncounter.id}
                  initial={{ opacity: 0, y: 10 }}
                  animate={{ opacity: 1, y: 0 }}
                  exit={{ opacity: 0, y: -10 }}
                  className="relative w-full max-w-4xl aspect-video border border-zinc-800 overflow-hidden"
                >
                  <img src={currentEncounter.image} className="w-full h-full object-cover opacity-60 grayscale group-hover:grayscale-0 transition-all duration-700" />
                  <div className="absolute inset-0 bg-gradient-to-t from-black via-transparent to-transparent" />
                  
                  {/* CLICK TO ADVANCE HINT */}
                  <div className="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
                    <div className="bg-white/10 backdrop-blur-md border border-white/20 px-6 py-3 uppercase text-[10px] tracking-[0.3em] font-black">
                      Click to Advance
                    </div>
                  </div>

                  {/* DNA OVERLAY */}
                  <div className="absolute bottom-4 left-4 right-4 md:bottom-8 md:left-8 md:right-8">
                    <div className="text-[10px] text-pink-500 font-bold mb-2 tracking-[0.3em] uppercase">Visual DNA Synthesis</div>
                    <div className="text-[10px] md:text-xs text-zinc-400 font-mono bg-black/60 p-2 md:p-3 border-l-2 border-pink-500 backdrop-blur-sm">
                      <span className="text-zinc-500">DNA_KEYWORDS:</span> {currentEncounter.visualPrompt.toUpperCase()}
                    </div>
                  </div>

              {/* ENCOUNTER INFO */}
                  <div className="absolute top-4 left-4 md:top-8 md:left-8">
                    <div className="bg-white text-black px-2 py-0.5 md:px-3 md:py-1 text-[8px] md:text-[10px] font-black uppercase mb-1 md:mb-2 inline-block">
                      {currentEncounter.type}
                    </div>
                    <h3 className="text-xl md:text-3xl font-black italic tracking-tighter">{currentEncounter.title}</h3>
                    <p className="text-zinc-400 max-w-xs md:max-w-md mt-1 md:mt-2 text-[10px] md:text-sm leading-relaxed">{currentEncounter.description}</p>
                  </div>
                </motion.div>
              )}
            </AnimatePresence>
          </div>

          {/* TIMELINE STRING */}
          <div className="p-4 md:p-8 border-t border-zinc-900 bg-zinc-950/50 backdrop-blur-xl">
            <div className="flex items-center gap-4 mb-6 overflow-x-auto pb-4 no-scrollbar">
              {encounters.map((e, i) => (
                <div key={`timeline-${e.id}-${i}`} className="flex items-center gap-4 shrink-0">
                  <div className={`w-3 h-3 rounded-full ${
                    e.type === 'COMBAT' ? 'bg-red-500' : 
                    e.type === 'RESOURCE' ? 'bg-blue-500' : 
                    'bg-white'
                  } ${i === encounters.length - 1 ? 'ring-4 ring-pink-500/20 scale-125' : 'opacity-40'}`} />
                  {i < encounters.length - 1 && <div className="w-12 h-[1px] bg-zinc-800" />}
                </div>
              ))}
              <div className="w-12 h-[1px] bg-zinc-800 border-dashed border-t" />
              <div className="w-3 h-3 rounded-full border border-zinc-800" />
            </div>

            <div className="flex justify-between items-center">
              <div className="flex flex-col">
                <div className="text-[10px] text-zinc-600 uppercase tracking-widest">
                  Step {encounters.length} / Narrative Timeline
                </div>
                {isTransitioning && (
                  <div className="text-[10px] text-red-500 font-bold animate-pulse uppercase mt-1">
                    Target Lock Detected - Transitioning to Combat...
                  </div>
                )}
              </div>
              <button 
                onClick={() => {
                  if (currentEncounter?.type === 'COMBAT') {
                    setGameState('ENCOUNTER');
                    setIsTransitioning(false);
                  } else {
                    advanceTimeline();
                  }
                }}
                disabled={o2 <= 0 || isTransitioning}
                className="bg-white text-black px-12 py-4 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all disabled:opacity-20 relative group"
              >
                <div className="flex flex-col items-center">
                  <span>{currentEncounter?.type === 'COMBAT' ? 'Enter Combat' : (o2 < 30 ? 'Desperate Move' : 'Advance Timeline')}</span>
                  {currentEncounter?.type !== 'COMBAT' && (
                    <span className="text-[8px] font-bold opacity-60 group-hover:opacity-100">
                      COST: 15 O2 | 5 FUEL
                    </span>
                  )}
                </div>
              </button>
            </div>
          </div>
        </div>
      )}

      {/* --- ENCOUNTER STATE (MOCK COMBAT) --- */}
      {gameState === 'ENCOUNTER' && (
        <div className="h-screen bg-red-950/10 flex items-center justify-center p-12">
          <motion.div 
            initial={{ opacity: 0, scale: 0.9 }}
            animate={{ opacity: 1, scale: 1 }}
            className="border-2 border-red-500/50 p-12 bg-black max-w-4xl w-full relative overflow-hidden"
          >
            <div className="absolute top-0 left-0 w-full h-1 bg-red-500 animate-pulse" />
            <h2 className="text-5xl font-black italic text-red-500 mb-4 tracking-tighter">COMBAT ENGAGED</h2>
            <p className="text-zinc-400 mb-12 text-lg">Iron Syndicate Scavengers are locking onto your position. Defensive systems active.</p>
            
            <div className="grid grid-cols-2 gap-8 mb-12">
              <div className="border border-zinc-800 p-6">
                <div className="text-xs text-zinc-500 mb-2 uppercase">
                  {selectedVehicle ? `Your ${selectedVehicle.type}` : 'Pilot Status'}
                </div>
                <div className={`text-2xl font-bold ${!selectedVehicle ? 'text-pink-500' : ''}`}>
                  {selectedVehicle ? selectedVehicle.class : 'UNARMORED PILOT'}
                </div>
                <div className="w-full h-2 bg-zinc-900 mt-4">
                  <motion.div 
                    initial={{ width: 0 }}
                    animate={{ width: selectedVehicle ? '85%' : '10%' }}
                    className={`h-full ${selectedVehicle ? 'bg-green-500' : 'bg-red-500 animate-pulse'}`} 
                  />
                </div>
                {!selectedVehicle && (
                  <div className="text-[8px] text-red-500 mt-2 font-bold uppercase">
                    Critical Risk: No Armor Detected
                  </div>
                )}
              </div>
              <div className="border border-zinc-800 p-6">
                <div className="text-xs text-zinc-500 mb-2 uppercase">Enemy</div>
                <div className="text-2xl font-bold text-red-500">SCAV-DRONE</div>
                <div className="w-full h-2 bg-zinc-900 mt-4">
                  <div className="h-full bg-red-500 w-[40%]" />
                </div>
              </div>
            </div>

            <button 
              onClick={finishEncounter}
              className={`w-full py-4 font-black uppercase tracking-tighter transition-all ${
                selectedVehicle 
                  ? 'bg-red-500 text-white hover:bg-white hover:text-black' 
                  : 'bg-zinc-800 text-zinc-500 hover:bg-pink-500 hover:text-white'
              }`}
            >
              {selectedVehicle ? 'Execute Final Strike' : 'Attempt Desperate Escape'}
            </button>
          </motion.div>
        </div>
      )}

      {/* --- DEBRIEF STATE --- */}
      {gameState === 'DEBRIEF' && (
        <div className="h-screen flex flex-col items-center justify-center p-8">
          <motion.div 
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            className="max-w-2xl w-full border border-zinc-800 p-12 bg-zinc-900/10 backdrop-blur-md"
          >
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
                <span className="text-zinc-500 uppercase text-xs">Resources Salvaged</span>
                <span className="font-bold text-blue-500">{encounters.filter(e => e.type === 'RESOURCE').length}</span>
              </div>
            </div>

            <button 
              onClick={() => window.location.reload()}
              className="w-full border border-white py-4 font-black uppercase tracking-tighter hover:bg-white hover:text-black transition-all"
            >
              Return to Mothership
            </button>
          </motion.div>
        </div>
      )}

      {/* SCANLINES EFFECT */}
      <div className="fixed inset-0 pointer-events-none opacity-[0.03] bg-[linear-gradient(rgba(18,16,16,0)_50%,rgba(0,0,0,0.25)_50%),linear-gradient(90deg,rgba(255,0,0,0.06),rgba(0,255,0,0.02),rgba(0,0,255,0.06))] z-50 bg-[length:100%_2px,3px_100%]" />
    </div>
  );
}
