import { Box, Button, Input, TextField, Typography } from '@mui/material'
import { Container } from '@mui/system'
import { useState } from 'react'
import reactLogo from './assets/react.svg'

function App() {
  const [stats, setStats] = useState()
  const [url, setUrl] = useState<string>()

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUrl(event.target.value);
  };

  const fetchStats = async () => {
    const res = await fetch('http://localhost:8080/player', {
      method: "POST",
      body: JSON.stringify({ url })
    })
    const jsonRes = await res.json()
    setStats(jsonRes)
  }

  return (
    <Container>
      <Typography variant="body1">Enter a player's baseball reference URL to get started</Typography>
      <Box display="flex">
        <TextField defaultValue={url} label="Baseball Reference URL" value={url} onChange={handleChange} />
        <Button onClick={fetchStats}>Search</Button>
      </Box>
      {JSON.stringify(stats)}
    </Container>
  )
}

export default App
