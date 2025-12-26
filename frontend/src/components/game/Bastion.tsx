'use client';

import ShowcaseEngine from '@/components/game/ShowcaseEngine';
import { useAuthSync } from '@/hooks/use-auth-sync';
import { useEffect, useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { LoginButton } from '@/components/login-button';
import { bastionSystem, BastionState } from '@/systems/BastionSystem';
import { gameEvents, GAME_EVENTS } from '@/systems/EventBus';

interface BastionProps {
  onDeploy: () => void;
  onGacha: () => void;
}

export default function Bastion({ onDeploy, onGacha }: BastionProps) {
  const { user: backendUser } = useAuthSync();
  const [bastionState, setBastionState] = useState<BastionState>(bastionSystem.getState());
  const [showInventory, setShowInventory] = useState(false);

  useEffect(() => {
    const unsubscribe = gameEvents.on(GAME_EVENTS.BASTION_UPDATED, (newState: BastionState) => {
      setBastionState(newState);
    });

    if (backendUser?.active_character_id || backendUser?.id) {
      bastionSystem.refreshVehicles(backendUser.active_character_id || backendUser.id);
    }

    return () => unsubscribe();
  }, [backendUser?.id, backendUser?.active_character_id]);

  const vehicles = bastionState.vehicles || [];
  const currentVehicle = vehicles.find(v => v.id === bastionState.selectedVehicleId) || vehicles[0];
  const equippedItems = bastionState.items.filter(item => item.is_equipped && item.parent_item_id === currentVehicle?.id);

  return (
    <div className="relative w-full h-screen bg-black overflow-hidden">
      <ShowcaseEngine vehicleId={currentVehicle?.id || "VHC-001-ALPHA"} />
      
      {/* UI Overlay */}
      <div className="absolute top-8 left-8 z-10">
        <div className="flex items-start gap-4">
          {/* Pilot ID Badge */}
          <div className="bg-zinc-900/80 border border-zinc-700 p-4 rounded-lg backdrop-blur-md shadow-xl flex gap-4 items-center">
            <div className="w-16 h-16 bg-zinc-800 border border-zinc-700 rounded overflow-hidden flex items-center justify-center">
              {backendUser?.active_character?.gender ? (
                <img 
                  src={`/images/pilots/${backendUser.active_character.gender.toLowerCase()}.png`} 
                  alt="Pilot Avatar"
                  className="w-full h-full object-cover"
                />
              ) : (
                <div className="text-pink-500 text-2xl font-bold">
                  {backendUser?.active_character?.name?.[0] || 'P'}
                </div>
              )}
            </div>
            <div>
              <h2 className="text-xl font-bold text-white tracking-tighter uppercase italic">
                {backendUser?.active_character?.name || 'UNREGISTERED PILOT'}
              </h2>
              <div className="flex gap-2 mt-1">
                <span className="text-[10px] font-mono text-zinc-400 bg-zinc-800 px-1.5 py-0.5 rounded">
                  {backendUser?.active_character?.gender || 'UNKNOWN'}
                </span>
                <span className="text-[10px] font-mono text-pink-400 bg-pink-400/10 border border-pink-400/20 px-1.5 py-0.5 rounded">
                  RANK {bastionState.pilotStats?.rank || 1}
                </span>
                <span className="text-[10px] font-mono text-yellow-400 bg-yellow-400/10 border border-yellow-400/20 px-1.5 py-0.5 rounded">
                  SCRAP: {bastionState.pilotStats?.scrap_metal || 0}
                </span>
                <span className="text-[10px] font-mono text-cyan-400 bg-cyan-400/10 border border-cyan-400/20 px-1.5 py-0.5 rounded">
                  DATA: {bastionState.pilotStats?.research_data || 0}
                </span>
              </div>
            </div>
          </div>

          {/* Bastion Global Modules */}
          <div className="bg-zinc-900/80 border border-zinc-700 p-4 rounded-lg backdrop-blur-md shadow-xl flex flex-col gap-2">
            <div className="text-[10px] font-bold text-zinc-500 uppercase tracking-widest border-b border-zinc-800 pb-1 mb-1">
              Bastion Global Modules
            </div>
            <div className="flex gap-4">
              <div className="flex flex-col items-center">
                <div className="text-[8px] text-zinc-400 uppercase">Radar</div>
                <div className="text-lg font-black text-cyan-400">LV.{bastionState.pilotStats?.metadata?.radar_level || 1}</div>
              </div>
              <div className="flex flex-col items-center">
                <div className="text-[8px] text-zinc-400 uppercase">Lab</div>
                <div className="text-lg font-black text-purple-400">LV.{bastionState.pilotStats?.metadata?.lab_level || 1}</div>
              </div>
              <div className="flex flex-col items-center">
                <div className="text-[8px] text-zinc-400 uppercase">Warp</div>
                <div className="text-lg font-black text-orange-400">LV.{bastionState.pilotStats?.metadata?.warp_level || 1}</div>
              </div>
            </div>
          </div>
        </div>
        
        <div className="mt-4 flex flex-col gap-2">
          <LoginButton />

          {backendUser?.wallet_address && (
            <div className="text-[10px] text-green-500 font-mono bg-green-500/10 border border-green-500/20 px-2 py-1 rounded">
              WALLET: {backendUser.wallet_address.slice(0, 6)}...{backendUser.wallet_address.slice(-4)}
            </div>
          )}
        </div>
        {currentVehicle && (
          <div className="mt-4 flex flex-col gap-2">
            {/* Vehicle Selector */}
            {vehicles.length > 1 && (
              <div className="flex gap-1 mb-2 overflow-x-auto pb-2 max-w-[300px]">
                {vehicles.map(v => (
                  <button
                    key={v.id}
                    onClick={() => bastionSystem.selectVehicle(v.id)}
                    className={`px-2 py-1 text-[8px] font-bold uppercase border transition-all whitespace-nowrap ${
                      v.id === currentVehicle.id 
                        ? 'bg-white text-black border-white' 
                        : 'bg-black/40 text-zinc-500 border-zinc-800 hover:border-zinc-600'
                    }`}
                  >
                    {v.class}
                  </button>
                ))}
              </div>
            )}
            
            <div className="flex gap-2">
              {currentVehicle.is_void_touched && (
                <span key="void-touched" className="px-2 py-1 bg-purple-900/50 border border-purple-500 text-purple-300 text-[10px] font-bold uppercase tracking-widest">
                  Void-Touched
                </span>
              )}
              <span key="tier" className="px-2 py-1 bg-zinc-900 border border-zinc-700 text-zinc-400 text-[10px] font-bold uppercase tracking-widest">
                Tier {currentVehicle.tier || 1}
              </span>
              <span key="rarity" className="px-2 py-1 bg-zinc-900 border border-zinc-700 text-zinc-400 text-[10px] font-bold uppercase tracking-widest">
                {currentVehicle.rarity}
              </span>
            </div>
            
            {/* CP Display */}
            <div className="bg-black/60 border-l-4 border-pink-500 p-3 backdrop-blur-sm">
              <div className="text-[10px] text-zinc-500 uppercase tracking-widest">Combat Power</div>
              <div className="text-3xl font-black text-white italic tracking-tighter">
                {bastionState.selectedVehicleCP} <span className="text-xs text-pink-500 not-italic ml-1">CP</span>
              </div>
            </div>
          </div>
        )}
      </div>

      {/* Inventory Toggle Button */}
      <div className="absolute top-8 right-8 z-10">
        <button 
          onClick={() => setShowInventory(!showInventory)}
          className={`p-4 border transition-all ${showInventory ? 'bg-white text-black border-white' : 'bg-black/40 text-white border-zinc-700 hover:border-white'}`}
        >
          <div className="text-[10px] font-bold uppercase tracking-widest">Vehicle Systems</div>
        </button>
      </div>

      {/* Inventory Overlay */}
      <AnimatePresence>
        {showInventory && (
          <motion.div 
            initial={{ x: 400, opacity: 0 }}
            animate={{ x: 0, opacity: 1 }}
            exit={{ x: 400, opacity: 0 }}
            className="absolute top-0 right-0 w-[400px] h-full bg-zinc-950/90 border-l border-zinc-800 backdrop-blur-xl z-20 p-8 overflow-y-auto"
          >
            <div className="flex justify-between items-center mb-8">
              <h3 className="text-xl font-black italic text-white uppercase tracking-tighter">System Inventory</h3>
              <button onClick={() => setShowInventory(false)} className="text-zinc-500 hover:text-white">CLOSE</button>
            </div>

            <div className="space-y-6">
              {/* Vehicle Stats Detail */}
              <div className="bg-black/40 border border-zinc-800 p-4">
                <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-4">Core Specifications</div>
                <div className="grid grid-cols-2 gap-4">
                  <div>
                    <div className="text-[10px] text-zinc-600 uppercase">Armor (HP)</div>
                    <div className="text-lg font-bold text-white">{currentVehicle?.stats.hp}</div>
                  </div>
                  <div>
                    <div className="text-[10px] text-zinc-600 uppercase">Firepower</div>
                    <div className="text-lg font-bold text-white">{currentVehicle?.stats.attack}</div>
                  </div>
                  <div>
                    <div className="text-[10px] text-zinc-600 uppercase">Shielding</div>
                    <div className="text-lg font-bold text-white">{currentVehicle?.stats.defense}</div>
                  </div>
                  <div>
                    <div className="text-[10px] text-zinc-600 uppercase">Agility</div>
                    <div className="text-lg font-bold text-white">{currentVehicle?.stats.speed}</div>
                  </div>
                </div>
              </div>

              {/* Equipped Parts - Visual Map */}
              <div>
                <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-4">Equipped Modules (Visual Map)</div>
                <div className="relative w-full aspect-[3/4] bg-black/40 border border-zinc-800 overflow-hidden mb-6">
                  {/* Silhouette Background */}
                  <div className="absolute inset-0 flex items-center justify-center opacity-10 pointer-events-none">
                    <svg viewBox="0 0 100 100" className="h-full w-auto text-white fill-current">
                      <path d="M50,5 C55,5 58,10 58,15 C58,20 55,25 50,25 C45,25 42,20 42,15 C42,10 45,5 50,5 Z M40,28 L60,28 C65,28 68,32 68,37 L68,55 C68,58 65,60 62,60 L58,60 L58,95 C58,98 55,100 52,100 L48,100 C45,100 42,98 42,95 L42,60 L38,60 C35,60 32,58 32,55 L32,37 C32,32 35,28 40,28 Z" />
                    </svg>
                  </div>

                  {/* Slots */}
                  {[
                    { id: 'HEAD', label: 'Head', top: '10%', left: '50%' },
                    { id: 'CORE', label: 'Core', top: '35%', left: '50%' },
                    { id: 'ARM_L', label: 'Left Arm', top: '35%', left: '20%' },
                    { id: 'ARM_R', label: 'Right Arm', top: '35%', left: '80%' },
                    { id: 'LEGS', label: 'Legs', top: '75%', left: '50%' },
                  ].map(slot => {
                    const item = equippedItems.find(i => i.slot === slot.id);
                    return (
                      <div 
                        key={slot.id}
                        className="absolute -translate-x-1/2 -translate-y-1/2 group"
                        style={{ top: slot.top, left: slot.left }}
                      >
                        <div className={`w-16 h-16 border-2 transition-all flex flex-col items-center justify-center p-1 text-center
                          ${item ? 'border-pink-500 bg-pink-500/10' : 'border-zinc-800 border-dashed bg-zinc-900/20'}
                          hover:border-white cursor-pointer
                        `}>
                          {item ? (
                            <>
                              <div className="text-[8px] font-black text-white truncate w-full">{item.name}</div>
                              <div className="text-[6px] text-pink-400 mt-1">{item.rarity}</div>
                              <button 
                                onClick={(e) => {
                                  e.stopPropagation();
                                  bastionSystem.unequipItem(item.id);
                                }}
                                className="absolute -top-2 -right-2 w-5 h-5 bg-red-500 text-white rounded-full flex items-center justify-center text-[10px] opacity-0 group-hover:opacity-100 transition-opacity"
                              >
                                ×
                              </button>
                            </>
                          ) : (
                            <>
                              <div className="text-[8px] text-zinc-600 font-bold uppercase">{slot.label}</div>
                              <div className="text-[6px] text-zinc-700 mt-1">EMPTY</div>
                            </>
                          )}
                        </div>
                        {/* Connector Line (Optional) */}
                        <div className="absolute top-1/2 left-1/2 -z-10 w-px h-px bg-zinc-800" />
                      </div>
                    );
                  })}
                </div>
              </div>

              {/* Equipped Parts List (Legacy View) */}
              <div>
                <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-4">Module Details</div>
                <div className="space-y-2">
                  {equippedItems.length > 0 ? equippedItems.map(item => (
                    <div key={item.id} className="bg-zinc-900/50 border border-zinc-800 p-3 flex justify-between items-center group hover:border-pink-500/50 transition-colors">
                      <div>
                        <div className="text-xs font-bold text-white uppercase">{item.name}</div>
                        <div className="text-[8px] text-zinc-500 uppercase">{item.slot || 'GENERAL SLOT'} // {item.rarity}</div>
                      </div>
                      <div className="text-right">
                        <div className="text-[10px] text-pink-500 font-mono">
                          {item.stats.bonus_attack ? `+${item.stats.bonus_attack} ATK` : ''}
                          {item.stats.bonus_defense ? `+${item.stats.bonus_defense} DEF` : ''}
                        </div>
                        <div className="w-16 h-1 bg-zinc-800 mt-1 mb-2">
                          <div className="h-full bg-green-500" style={{ width: `${(item.durability / item.max_durability) * 100}%` }} />
                        </div>
                        <button 
                          onClick={() => bastionSystem.unequipItem(item.id)}
                          className="text-[8px] text-zinc-500 hover:text-white uppercase tracking-widest border border-zinc-800 px-2 py-0.5 hover:border-zinc-500 transition-all"
                        >
                          Unequip
                        </button>
                      </div>
                    </div>
                  )) : (
                    <div className="text-[10px] text-zinc-600 italic p-4 border border-dashed border-zinc-800 text-center">
                      No modules equipped. Visit the Gacha to find parts.
                    </div>
                  )}
                </div>
              </div>

              {/* Unequipped Items */}
              <div>
                <div className="text-[10px] text-zinc-500 uppercase tracking-widest mb-4">Storage Bay</div>
                <div className="grid grid-cols-1 gap-2">
                  {bastionState.items.filter(i => !i.is_equipped).map(item => (
                    <div key={item.id} className="bg-black/20 border border-zinc-900 p-3 flex justify-between items-center opacity-60 hover:opacity-100 transition-opacity">
                      <div>
                        <div className="text-xs font-bold text-zinc-400 uppercase">{item.name}</div>
                        <div className="text-[8px] text-zinc-600 uppercase">{item.item_type}</div>
                      </div>
                      <button 
                        onClick={() => bastionSystem.equipItem(item.id)}
                        className="text-[8px] border border-zinc-700 px-2 py-1 hover:bg-white hover:text-black transition-colors"
                      >
                        EQUIP
                      </button>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          </motion.div>
        )}
      </AnimatePresence>

      <div className="absolute bottom-8 right-8 z-10 flex flex-col gap-4">
        <button 
          onClick={onGacha}
          className="bg-yellow-400 text-black px-6 py-2 font-bold uppercase tracking-tighter hover:bg-yellow-300 transition-colors text-center"
        >
          Void Signals (Gacha)
        </button>
        <button className="bg-white text-black px-6 py-2 font-bold uppercase tracking-tighter hover:bg-gray-200 transition-colors">
          Share to Social
        </button>
        <button 
          onClick={onDeploy}
          className="border border-white text-white px-6 py-2 font-bold uppercase tracking-tighter hover:bg-white hover:text-black transition-colors text-center"
        >
          Launch Expedition
        </button>
      </div>
    </div>
  );
}
