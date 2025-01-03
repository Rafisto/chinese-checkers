import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { GlobalStateProvider } from './hooks/globalState';
import Board from './board'
import Menu from './menu'
import './main.css';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <GlobalStateProvider>
      <div className="grid">
        <Board />
        <Menu />
      </div>
    </GlobalStateProvider>
  </StrictMode>,
)
