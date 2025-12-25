'use client';

import { motion } from 'framer-motion';
import { Sector } from '@/services/exploration';

interface UniverseMapProps {
  sectors: Sector[];
  selectedSector: Sector | null;
  onSelectSector: (sector: Sector) => void;
  onScanSector: () => void;
  onBack: () => void;
  selectedVehicleClass?: string;
}

export default function UniverseMap({ 
  sectors, 
  selectedSector, 
  onSelectSector, 
  onScanSector, 
  onBack,
  selectedVehicleClass 
}: UniverseMapProps) {
  return (
    <div className="h-screen flex flex-col p-8 bg-black text-white font-mono">
      <div className="flex justify-between items-center mb-8">
        <div>
          <h2 className="text-3xl font-black italic tracking-tighter text-pink-500">UNIVERSE MAP</h2>
          <p className="text-zinc-500 text-xs uppercase tracking-[0.3em]">Select Destination Sector</p>
        </div>
        <div className="text-right">
          <div className="text-[10px] text-zinc-500 uppercase mb-1">Selected Vehicle</div>
          <div className="text-sm font-bold italic">{selectedVehicleClass || 'PILOT'}</div>
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
              onClick={() => onSelectSector(s)}
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
            onClick={onScanSector}
            disabled={!selectedSector}
            className="w-full bg-white text-black py-4 font-black uppercase tracking-tighter hover:bg-pink-500 hover:text-white transition-all disabled:opacity-20"
          >
            Scan Sector
          </button>
          <button 
            onClick={onBack}
            className="w-full border border-zinc-800 text-zinc-500 py-2 text-xs uppercase tracking-widest hover:text-white transition-all"
          >
            Return to Bastion
          </button>
        </div>
      </div>
    </div>
  );
}
