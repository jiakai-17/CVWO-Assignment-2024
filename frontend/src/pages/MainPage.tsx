import { ChangeEvent, useContext, useEffect, useState } from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { CircularProgress, Divider, List, ListItem, Pagination, Stack, useMediaQuery } from "@mui/material";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import SearchButton from "../components/SearchButton.tsx";
import SortButton from "../components/SortButton.tsx";
import { Link, useNavigate, useSearchParams } from "react-router-dom";
import ThreadPreview from "../components/ThreadPreview.tsx";
import AuthContext from "../contexts/AuthContext.tsx";
import Thread from "../models/thread/Thread.tsx";

export default function Page() {
  const [isLogin, setIsLogin] = useState(false);
  const navigate = useNavigate();
  const { auth } = useContext(AuthContext);

  useEffect(() => {
    setIsLogin(auth.isLogin);
  }, [auth.isLogin]);

  const isSmallScreen = useMediaQuery("(max-width: 600px)");

  // Searching for threads
  // Get query from URL, if any
  const [threads, setThreads] = useState([] as Thread[]);
  const [searchParams] = useSearchParams();
  const hasQuery = searchParams.has("q");
  const [searchQuery, setSearchQuery] = useState(hasQuery ? searchParams.get("q") : "");

  // for the search box
  const [inputQuery, setInputQuery] = useState("searchQuery");

  // for pagination
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(10);
  const handlePageChange = (_event: ChangeEvent<unknown>, value: number) => {
    console.log("page changed", value);
    setCurrentPage(value);
  };

  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    console.log("input query updated", inputQuery);
  }, [inputQuery]);

  useEffect(() => {
    if (hasQuery) {
      console.log("found query in url", searchQuery);
      setSearchQuery(searchParams.get("q") ?? "");
    }
  }, [searchQuery, searchParams, hasQuery]);

  const handleSearch = () => {
    console.log("search button clicked", inputQuery);
    setSearchQuery(inputQuery);
    navigate("/?q=" + inputQuery);
  };

  // Sorting threads by criteria
  const [threadSortCriteria, setThreadSortCriteria] = useState("Newest first");

  const availableThreadSortCriteria = new Map<string, string>([
    ["Newest first", "created_time_desc"],
    ["Oldest first", "created_time_asc"],
    ["Most comments first", "num_comments_desc"],
    ["Least comments first", "num_comments_asc"],
  ]);

  useEffect(() => {
    console.log(threadSortCriteria);
  }, [threadSortCriteria]);

  useEffect(() => {
    setInputQuery(searchQuery ?? "");
    console.log("search query updated", searchQuery);
    setIsLoading(true);
    fetch(
      "/api/v1/thread?q=" +
        searchQuery +
        "&p=" +
        currentPage +
        "&order=" +
        availableThreadSortCriteria.get(threadSortCriteria),
    )
      .then((response) => {
        if (response.ok) {
          return response.json();
        } else {
          throw new Error("Something went wrong");
        }
      })
      .then((data) => {
        console.log(data);
        setThreads(data);
        setIsLoading(false);
      })
      .catch((error) => {
        console.log(error);
        setIsLoading(false);
      });
  }, [currentPage, searchQuery, threadSortCriteria]);

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
          onKeyDown={(event) => {
            if (event.key === "Enter") {
              handleSearch();
            }
          }}
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
      {!isLoading && (
        <Stack alignItems="center">
          <Pagination
            count={totalPages}
            page={currentPage}
            onChange={handlePageChange}
          />
          {threads.length === 0 && (
            <Typography
              variant="h6"
              className={"px-6 py-20"}
            >
              No threads found
            </Typography>
          )}
          {threads.length > 0 && (
            <List sx={{ width: "100%", bgcolor: "background.paper" }}>
              {threads.map((thread) => (
                <ListItem key={thread.id}>
                  <ThreadPreview {...thread} />
                </ListItem>
              ))}
            </List>
          )}
          <Pagination
            count={totalPages}
            page={currentPage}
            onChange={handlePageChange}
          />
        </Stack>
      )}
      {isLoading && (
        <Stack alignItems={"center"}>
          <div className={"flex flex-row"}>
            <CircularProgress />
            <Typography
              variant="h6"
              className={"px-6"}
            >
              Loading threads...
            </Typography>
          </div>
        </Stack>
      )}
    </Box>
  );
}
