"use client";

import * as React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Divider, List, ListItem } from "@mui/material";
import ThreadComponent from "@/components/ThreadComponent";
import sampleThreads from "@/models/thread/SampleThreads";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import SearchIcon from "@mui/icons-material/Search";
import SortIcon from "@mui/icons-material/Sort";

export default function Page() {
  return (
    <Box sx={{ mx: "2", textAlign: "center", mb: 10 }}>
      <Typography variant="h4" sx={{ px: 2 }}>Welcome to CVWO Forum</Typography>
      <Divider sx={{ m: 5 }} />
      <Box sx={{ mx: 4, mb: 4, display: "flex", alignItems: "stretch", width: "full", gap: 2 }}>
        <TextField
          id="filled-search"
          label="Search for a thread..."
          type="search"
          variant="filled"
          sx={{ flexGrow: 1 }}
        />
        <Button
          variant="contained"
          disableElevation
          size="large"
          endIcon={<SearchIcon />}>
          Search
        </Button>
        <Button
          variant="outlined"
          size="large"
          endIcon={<SortIcon />}>
          Sort...
        </Button>
      </Box>
      <List sx={{ width: "100%", bgcolor: "background.paper" }}>
        {sampleThreads.map((thread) => (
          <ListItem key={thread.id}>
            <ThreadComponent {...thread} />
          </ListItem>
        ))}
      </List>
    </Box>

  );
}
