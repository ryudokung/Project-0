'use client';

import { motion } from 'framer-motion';
import { SubSector, PlanetLocation } from '@/services/exploration';

interface PlanetSurfaceProps {
  selectedSubSector: SubSector;
  selectedPlanetLocation: PlanetLocation | null;
  onSelectPlanetLocation: (loc: PlanetLocation) => void;
  onConfirmDeployment: () => void;
  onBack: () => void;
  vehicles: any[];
  selectedVehicle: any | null;
  onSelectVehicle: (v: any | null) => void;
  inventory: string[];
  canDeploy: (target: any) => boolean;
  meetsRequirements: (reqs: string[]) => boolean;
}

export default function PlanetSurface({
  selectedSubSector,
  selectedPlanetLocation,
  onSelectPlanetLocation,
  onConfirmDeployment,
  onBack,
  vehicles,
  selectedVehicle,
  onSelectVehicle,
  inventory,
  canDeploy,
  meetsRequirements
}: PlanetSurfaceProps) {
  return (
    <div className="h-screen flex flex-col p-8 bg-black text-white font-mono">
      <div className="flex justify-between items-center mb-8">
        <div>
          <h2 className="text-3xl font-black italic tracking-tighter text-pink-500">{selectedSubSector.name}</h2>
          <p className="text-zinc-500 text-xs uppercase tracking-[0.3em]">Surface Scan / Select Tactical Objective</p>
        </div>
        <button 
          onClick={onBack}
          className="text-zinc-500 hover:text-white text-xs uppercase tracking-widest"
        >
          Back to Sector Scan
        </button>
      </div>

      <div className="flex-1 grid md:grid-cols-3 gap-8">
        {/* Surface Map Visualization */}
        <div className="md:col-span-2 border border-zinc-800 bg-zinc-900/10 relative overflow-hidden">
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
          
          {selectedSubSector.locations?.map((loc, idx) => (
            <motion.div
              key={`location-${loc.id || idx}`}
              initial={{ scale: 0 }}
              animate={{ scale: 1 }}
              whileHover={{ scale: 1.2 }}
              onClick={() => onSelectPlanetLocation(loc)}
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

                <div className="space-y-3">
                  <div className="text-[8px] text-zinc-500 uppercase">Select Deployment Asset</div>
                  <div className="grid grid-cols-1 gap-2">
                    <div 
                      onClick={() => onSelectVehicle(null)}
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
                        onClick={() => onSelectVehicle(v)}
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
            onClick={onConfirmDeployment}
            disabled={!selectedPlanetLocation || !meetsRequirements(selectedPlanetLocation.requirements) || !canDeploy(selectedPlanetLocation)}
            className="w-full bg-white text-black py-4 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all disabled:opacity-20"
          >
            Initiate Deployment
          </button>
        </div>
      </div>
    </div>
  );
}
