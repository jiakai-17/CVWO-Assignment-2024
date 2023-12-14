"use client";

import * as React from "react";
import { useEffect, useState } from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Divider, List, ListItem, useMediaQuery } from "@mui/material";
import ThreadComponent from "@/components/ThreadComponent";
import sampleThreads from "@/models/thread/SampleThreads";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import Link from "next/link";
import SortButton from "@/components/SortButton";
import SearchButton from "@/components/SearchButton";
import { useRouter, useSearchParams } from "next/navigation";

export default function Page() {

  const [isLogin, setIsLogin] = useState(false);

  useEffect(() => {
    setIsLogin(JSON.parse(localStorage.getItem("isLogin") ?? "false"));
  }, [localStorage.getItem("isLogin")]);

  const isSmallScreen = useMediaQuery("(max-width: 600px)");

  // Searching for threads
  // Get query from URL, if any
  const router = useRouter();
  const urlParams = useSearchParams();
  const hasQuery = urlParams.has("q");
  const [searchQuery, setSearchQuery] =
    useState(urlParams.get("q") ?? "");

  const [inputQuery, setInputQuery] = useState(searchQuery);

  useEffect(() => {
    console.log("input query updated", inputQuery);
  }, [inputQuery]);

  useEffect(() => {
    setInputQuery(searchQuery);
    console.log("search query updated", searchQuery);
    console.log("{MOCK API CALL TIME}", "Searching for threads with query", searchQuery);
  }, [searchQuery]);

  useEffect(() => {
    setSearchQuery(urlParams.get("q") ?? "");
    console.log("found query in url", searchQuery);
  }, [searchQuery, urlParams]);

  const handleSearch = () => {
    console.log("search button clicked", inputQuery);
    router.push("/?q=" + inputQuery);
  };

  // Sorting threads by criteria
  const [threadSortCriteria, setThreadSortCriteria] = useState("Newest first");

  const availableThreadSortCriteria = new Map<string, string>([
    ["Newest first", "date desc"],
    ["Oldest first", "date asc"],
    ["Most comments first", "comments desc"],
    ["Least comments first", "comments asc"],
  ]);

  useEffect(() => {
    console.log(threadSortCriteria);
  }, [threadSortCriteria]);

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
          value={inputQuery}
          onChange={(event) => setInputQuery(event.target.value)}
        />
        {isSmallScreen &&
          <>
            <SearchButton
              size={"small"}
              onClick={handleSearch}
            />
            <SortButton
              availableSortCriteriaMappings={availableThreadSortCriteria}
              setSortCriteria={setThreadSortCriteria}
              size={"small"}
            />
          </>
        }
        {!isSmallScreen &&
          <>
            <SearchButton
              size={"large"}
              onClick={handleSearch}
            />
            <SortButton
              availableSortCriteriaMappings={availableThreadSortCriteria}
              setSortCriteria={setThreadSortCriteria}
              size={"large"}
            />
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
