'use client';

import { useEffect, useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { gameEvents, GAME_EVENTS } from '@/systems/EventBus';

export default function NotificationSystem() {
  const [notifications, setNotifications] = useState<any[]>([]);

  useEffect(() => {
    const unsubscribe = gameEvents.on(GAME_EVENTS.NOTIFICATION, (data) => {
      const id = Math.random().toString(36).substr(2, 9);
      setNotifications(prev => [...prev, { ...data, id }]);
      
      setTimeout(() => {
        setNotifications(prev => prev.filter(n => n.id !== id));
      }, 3000);
    });

    return () => unsubscribe();
  }, []);

  return (
    <div className="fixed top-8 right-8 z-[100] flex flex-col gap-2 pointer-events-none">
      <AnimatePresence>
        {notifications.map(n => (
          <motion.div
            key={n.id}
            initial={{ opacity: 0, x: 50 }}
            animate={{ opacity: 1, x: 0 }}
            exit={{ opacity: 0, scale: 0.9 }}
            className={`px-6 py-3 border font-black italic text-sm tracking-tighter backdrop-blur-md ${
              n.type === 'ERROR' || n.type === 'CRITICAL' 
                ? 'bg-red-500/20 border-red-500 text-red-500' 
                : 'bg-white/10 border-white/20 text-white'
            }`}
          >
            {n.message}
          </motion.div>
        ))}
      </AnimatePresence>
    </div>
  );
}
