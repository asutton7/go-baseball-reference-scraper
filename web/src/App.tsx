import { Box, Button, Input, TextField, Typography } from '@mui/material'
import { Container } from '@mui/system'
import { useState } from 'react'
import { Line } from 'react-chartjs-2'
import { Chart, CategoryScale, LinearScale, PointElement, LineElement } from "chart.js";

Chart.register(CategoryScale, LinearScale, PointElement, LineElement);

function App() {
  const [stats, setStats] = useState<unknown[]>()
  const [url, setUrl] = useState<string>()

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUrl(event.target.value);
  };

  const fetchStats = async () => {
    const res = await fetch('http://localhost:8080/player?path=/players/w/willite01.shtml', {
      method: "POST",
      body: JSON.stringify({ url })
    })
    const jsonRes = await res.json()
    setStats(jsonRes)
  }

  const mapData = () => {
    if (stats) {
      return {
        datasets: [{
          data: stats.map((year: Record<string, number>) => ({
            x: year.Year,
            y: year.BA
          })),
          backgroundColor: 'black',
          borderColor: 'lightblue',
        }
        ]
      }
    }
  }

  return (
    <Container>
      <Typography variant="body1">Enter a player's baseball reference URL to get started</Typography>
      <Box display="flex">
        <TextField defaultValue={url} label="Baseball Reference URL" value={url} onChange={handleChange} />
        <Button onClick={fetchStats}>Search</Button>
      </Box>
      {stats && <Line id="line" data={mapData()} />}
    </Container>
  )
}

export default App
