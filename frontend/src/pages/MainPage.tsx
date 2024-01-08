import { useEffect, useState } from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Divider, List, ListItem, useMediaQuery } from "@mui/material";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import SearchButton from "../components/SearchButton.tsx";
import SortButton from "../components/SortButton.tsx";
import { Link, useNavigate, useSearchParams } from "react-router-dom";
import sampleThreads from "../models/thread/SampleThreads.tsx";
import ThreadPreview from "../components/ThreadPreview.tsx";

export default function Page() {
  const [isLogin, setIsLogin] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    setIsLogin(JSON.parse(localStorage.getItem("isLogin") ?? "false"));
  }, [localStorage.getItem("isLogin")]);

  const isSmallScreen = useMediaQuery("(max-width: 600px)");

  // Searching for threads
  // Get query from URL, if any
  const [searchParams] = useSearchParams();
  const hasQuery = searchParams.has("q");
  const [searchQuery, setSearchQuery] = useState(hasQuery ? searchParams.get("q") : "");

  // for the search box
  const [inputQuery, setInputQuery] = useState("searchQuery");

  useEffect(() => {
    console.log("input query updated", inputQuery);
  }, [inputQuery]);

  useEffect(() => {
    setInputQuery(searchQuery ?? "");
    console.log("search query updated", searchQuery);
    console.log("{MOCK API CALL TIME}", "Searching for threads with query", searchQuery);
  }, [searchQuery]);

  useEffect(() => {
    if (hasQuery) {
      console.log("found query in url", searchQuery);
      setSearchQuery(searchParams.get("q") ?? "");
      return;
    }
  }, [searchQuery, searchParams, hasQuery]);

  const handleSearch = () => {
    console.log("search button clicked", inputQuery);
    navigate("/?q=" + inputQuery);
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
    <Box className={"mx-2 mb-10 mt-16 text-center"}>
      <Typography
        variant="h4"
        className={"px-2"}
      >
        Welcome to CVWO Forum
      </Typography>
      <Divider sx={{ mx: 3, my: 6 }} />
      <Box className={"xs:mx-4 mx-6 mb-4 flex flex-row gap-4"}>
        <TextField
          id="filled-search"
          label="Search for a thread..."
          type="search"
          variant="filled"
          className={"flex-grow"}
          value={inputQuery}
          onChange={(event) => setInputQuery(event.target.value)}
        />
        {isSmallScreen && (
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
        )}
        {!isSmallScreen && (
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
        )}
      </Box>
      {isLogin && (
        <Box className={"xs:mx-4 mx-6 flex flex-row justify-end"}>
          <div />
          <Link to={"/new"}>
            <Button
              variant="contained"
              disableElevation
              size="large"
              sx={{ my: 2 }}
            >
              Create a new Thread
            </Button>
          </Link>
        </Box>
      )}
      <List sx={{ width: "100%", bgcolor: "background.paper" }}>
        {sampleThreads.map((thread) => (
          <ListItem key={thread.id}>
            <ThreadPreview {...thread} />
          </ListItem>
        ))}
      </List>
    </Box>
  );
}
