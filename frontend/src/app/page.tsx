"use client";

import * as React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Divider, List, ListItem, useMediaQuery } from "@mui/material";
import ThreadComponent from "@/components/ThreadComponent";
import sampleThreads from "@/models/thread/SampleThreads";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import SearchIcon from "@mui/icons-material/Search";
import SortIcon from "@mui/icons-material/Sort";
import { useEffect, useState } from "react";
import Link from "next/link";

export default function Page() {

  const [isLogin, setIsLogin] = useState(false);

  useEffect(() => {
    setIsLogin(JSON.parse(localStorage.getItem("isLogin") ?? "false"));
  }, [localStorage.getItem("isLogin")]);

  const isSmallScreen = useMediaQuery("(max-width: 600px)");

  return (
    <Box sx={{ mx: "2", textAlign: "center", mb: 10 }}>
      <Typography variant="h4" sx={{ px: 2 }}>Welcome to CVWO Forum</Typography>
      <Divider sx={{ m: 5 }} />
      <Box sx={{
        mx: { xs: 2, sm: 4 },
        mb: 2,
        display: "flex",
        alignItems: "stretch",
        width: "full",
        gap: 2,
      }}>
        <TextField
          id="filled-search"
          label="Search for a thread..."
          type="search"
          variant="filled"
          sx={{ flexGrow: 1 }}
        />
        {isSmallScreen &&
          <>
            <Button
              variant="contained"
              disableElevation
              size="small">
              <SearchIcon />
            </Button>
            <Button
              variant="outlined"
              size="small">
              <SortIcon />
            </Button>
          </>
        }
        {!isSmallScreen &&
          <>
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
          </>
        }
      </Box>
      {isLogin &&
        <Box sx={{
          mx: { xs: 2, sm: 4 },
          display: "flex",
          width: "full",
          justifyItems: "end",
        }}>
          <div style={{ flexGrow: 1 }}></div>
          <Link href={"/new"}>
            <Button
              variant="contained"
              disableElevation
              size="large"
              sx={{ my: 2 }}>
              Create a new Thread
            </Button>
          </Link>
        </Box>
      }
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
