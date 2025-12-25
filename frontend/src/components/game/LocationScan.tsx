'use client';

import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';
import { Sector, SubSector } from '@/services/exploration';
import { bastionSystem, BastionState } from '@/systems/BastionSystem';
import { gameEvents, GAME_EVENTS } from '@/systems/EventBus';

interface LocationScanProps {
  selectedSector: Sector;
  selectedSubSector: SubSector | null;
  onSelectSubSector: (ss: SubSector) => void;
  onConfirmDeployment: () => void;
  onBack: () => void;
  selectedVehicle: any | null;
  onSelectVehicle: (v: any | null) => void;
  upgrades: any;
  inventory: string[];
  canDeploy: (target: any) => boolean;
  meetsRequirements: (reqs: string[]) => boolean;
}

export default function LocationScan({
  selectedSector,
  selectedSubSector,
  onSelectSubSector,
  onConfirmDeployment,
  onBack,
  selectedVehicle,
  onSelectVehicle,
  upgrades,
  inventory,
  canDeploy,
  meetsRequirements
}: LocationScanProps) {
  const [bastionState, setBastionState] = useState<BastionState>(bastionSystem.getState());

  useEffect(() => {
    const unsubscribe = gameEvents.on(GAME_EVENTS.BASTION_UPDATED, (newState: BastionState) => {
      setBastionState(newState);
    });
    return () => unsubscribe();
  }, []);

  const vehicles = bastionState.vehicles;

  return (
    <div className="h-screen flex flex-col p-8 bg-black text-white font-mono">
      <div className="flex justify-between items-center mb-8">
        <div>
          <h2 className="text-3xl font-black italic tracking-tighter text-pink-500">{selectedSector.name}</h2>
          <p className="text-zinc-500 text-xs uppercase tracking-[0.3em]">Sector Scan / Identify Points of Interest</p>
        </div>
        <button 
          onClick={onBack}
          className="text-zinc-500 hover:text-white text-xs uppercase tracking-widest"
        >
          Back to Universe Map
        </button>
      </div>

      <div className="flex-1 grid md:grid-cols-3 gap-8">
        {/* Local Map Visualization */}
        <div className="md:col-span-2 border border-zinc-800 bg-zinc-900/10 relative overflow-hidden">
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
          
          {selectedSector.subSectors.map((ss, idx) => (
            <motion.div
              key={`subsector-${ss.id || idx}`}
              initial={{ scale: 0 }}
              animate={{ scale: 1 }}
              whileHover={{ scale: 1.2 }}
              onClick={() => onSelectSubSector(ss)}
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
                      <span>Vehicle Deployment</span>
                      <span className={selectedSubSector.suitability.vehicle > 50 ? 'text-green-500' : 'text-red-500'}>
                        {selectedSubSector.suitability.vehicle}%
                      </span>
                    </div>
                    <div className="w-full h-1 bg-zinc-900">
                      <div className="h-full bg-pink-500" style={{ width: `${selectedSubSector.suitability.vehicle}%` }} />
                    </div>
                  </div>
                </div>

                <div className="space-y-3">
                  <div className="text-[8px] text-zinc-500 uppercase">Select Deployment Asset</div>
                  <div className="grid grid-cols-1 gap-2">
                    <div 
                      onClick={() => onSelectVehicle(null)}
                      className={`p-2 border text-[10px] cursor-pointer transition-all ${!selectedVehicle ? 'border-white bg-white/10' : 'border-zinc-800 hover:border-zinc-700'}`}
                    >
                      <div className="flex justify-between">
                        <span className="font-bold">PILOT ONLY</span>
                        <span className={selectedSubSector.allowedModes.includes('PILOT') ? 'text-green-500' : 'text-red-500'}>
                          {selectedSubSector.allowedModes.includes('PILOT') ? '✓' : '✗'}
                        </span>
                      </div>
                    </div>
                    {vehicles.map((v, idx) => {
                      const vType = v.vehicle_type || v.type || 'VEHICLE';
                      const vClass = v.class || v.model || 'UNKNOWN';
                      return (
                        <div 
                          key={`vehicle-deploy-subsector-${v.id || idx}`}
                          onClick={() => onSelectVehicle(v)}
                          className={`p-2 border text-[10px] cursor-pointer transition-all ${selectedVehicle?.id === v.id ? 'border-pink-500 bg-pink-500/10' : 'border-zinc-800 hover:border-zinc-700'}`}
                        >
                          <div className="flex justify-between">
                            <span className="font-bold">{vClass} ({vType})</span>
                            <span className={selectedSubSector.allowedModes.includes(vType) ? 'text-green-500' : 'text-red-500'}>
                              {selectedSubSector.allowedModes.includes(vType) ? '✓' : '✗'}
                            </span>
                          </div>
                        </div>
                      );
                    })}
                  </div>
                </div>

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

                {!canDeploy(selectedSubSector) && (
                  <div className="p-3 border border-red-500 bg-red-500/10 text-[10px] text-red-500 font-bold uppercase">
                    Deployment Blocked: Check Compatibility
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
            onClick={onConfirmDeployment}
            disabled={!selectedSubSector || !meetsRequirements(selectedSubSector.requirements) || !canDeploy(selectedSubSector)}
            className="w-full bg-white text-black py-4 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all disabled:opacity-20"
          >
            {selectedSubSector?.type === 'PLANET' ? 'Scan Surface' : 'Initiate Deployment'}
          </button>
        </div>
      </div>
    </div>
  );
}
