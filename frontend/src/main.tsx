import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import Board from './board'
import Menu from './menu'
import './main.css';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <div className="grid">
      <Board/>
      <Menu/>
    </div>
  </StrictMode>,
)
