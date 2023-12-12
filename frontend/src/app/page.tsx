"use client";

import * as React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Divider, List, ListItem } from "@mui/material";
import ThreadComponent from "@/components/ThreadComponent";
import sampleThreads from "@/models/thread/SampleThreads";

export default function Page() {
  return (
    <Box sx={{ mx: "auto", textAlign: "center", mb: 10 }}>
      <Typography variant="h4" sx={{ px: 2 }}>Welcome to CVWO Forum</Typography>
      <Divider sx={{ m: 5 }} />
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
