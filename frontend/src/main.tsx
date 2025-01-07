import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { GlobalStateProvider } from './hooks/globalStateProvider';
import Grid from './grid';
import Game from './game'
import Menu from './menu'
import './styles/main.css';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <GlobalStateProvider>
      <Grid>
        <Game />
        <Menu /> 
      </Grid>
    </GlobalStateProvider>
  </StrictMode>,
)
