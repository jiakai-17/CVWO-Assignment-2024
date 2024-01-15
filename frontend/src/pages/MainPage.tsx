import { ChangeEvent, useCallback, useContext, useEffect, useMemo, useState } from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import {
  Alert,
  CircularProgress,
  Divider,
  IconButton,
  List,
  ListItem,
  Pagination,
  Stack,
  useMediaQuery,
} from "@mui/material";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import SearchButton from "../components/SearchButton.tsx";
import SortButton from "../components/SortButton.tsx";
import { Link, useNavigate, useSearchParams } from "react-router-dom";
import ThreadPreview from "../components/ThreadPreview.tsx";
import AuthContext from "../contexts/AuthContext.tsx";
import Thread from "../models/Thread.tsx";
import ClearIcon from "@mui/icons-material/Clear";

export default function Page() {
  const [isLogin, setIsLogin] = useState(false);
  const navigate = useNavigate();

  const [isError, setIsError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  // Update login status
  const { auth, isLoaded } = useContext(AuthContext);
  useEffect(() => {
    if (!isLoaded) {
      return;
    }
    setIsLogin(auth.isLogin);
  }, [auth.isLogin, isLoaded]);

  // Responsive design
  const isSmallScreen = useMediaQuery("(max-width: 600px)");

  // Search
  // Get the query from the URL
  const [searchParams] = useSearchParams();
  const hasSearchQuery = searchParams.has("q");

  // searchQuery is the query to be used in the API call
  const [searchQuery, setSearchQuery] = useState(hasSearchQuery ? searchParams.get("q") : "");

  // inputQuery is the query to be displayed in the search bar
  const [inputQuery, setInputQuery] = useState(hasSearchQuery ? searchParams.get("q") : "");

  // List of threads to display
  const [threads, setThreads] = useState([] as Thread[]);

  // Handle pagination
  const threadsPerPage = 10;
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0);
  const handlePageChange = (_event: ChangeEvent<unknown>, value: number) => {
    setCurrentPage(value);
  };

  // Sorting threads by criteria
  const [threadSortCriteria, setThreadSortCriteria] = useState("Newest first");
  const availableThreadSortCriteria = useMemo(
    () =>
      new Map<string, string>([
        ["Newest first", "created_time_desc"],
        ["Oldest first", "created_time_asc"],
        ["Most comments first", "num_comments_desc"],
        ["Least comments first", "num_comments_asc"],
      ]),
    [],
  );

  // API call to search
  const [isLoading, setIsLoading] = useState(true);
  const handleSearch = useCallback(() => {
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
          response.json().then((data) => {
            setThreads(data.threads);
            setTotalPages(Math.max(1, Math.ceil(data.total_threads / threadsPerPage)));
          });
        } else {
          response.text().then((text) => {
            setIsError(true);
            setErrorMessage(text);
            console.error(text);
          });
        }
      })
      .then(() => {
        setIsLoading(false);
      });
  }, [availableThreadSortCriteria, currentPage, searchQuery, threadSortCriteria]);

  // Updates the URL with the new search query and sort criteria
  const updateUrl = (searchQuery: string, sortCriteria: string) => {
    navigate("/?q=" + searchQuery + "&order=" + availableThreadSortCriteria.get(sortCriteria));
  };

  // Runs a search whenever the url changes
  useEffect(() => {
    const hasSearchQuery = searchParams.has("q");
    if (hasSearchQuery) {
      setSearchQuery(searchParams.get("q") ?? "");
      setInputQuery(searchParams.get("q") ?? "");
      setCurrentPage(1);
    }
    handleSearch();
  }, [handleSearch, hasSearchQuery, searchParams, searchQuery, setSearchQuery]);

  // Updates the URL when the sort criteria changes
  const handleSortCriteriaClick = (newCriteria: string) => {
    setThreadSortCriteria(newCriteria);
    setCurrentPage(1);
    updateUrl(searchQuery ?? "", newCriteria);
  };

  // Updates the URL when the search button is clicked
  const handleSearchButtonClick = () => {
    setSearchQuery(inputQuery);
    setCurrentPage(1);
    updateUrl(inputQuery ?? "", threadSortCriteria);
  };

  // Updates the URL when the clear button is clicked
  const resetSearch = () => {
    setInputQuery("");
    setThreadSortCriteria("Newest first");
    setSearchQuery("");
    setCurrentPage(1);
    navigate("/");
  };

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
              handleSearchButtonClick();
            }
          }}
          InputProps={{
            endAdornment: (
              <IconButton>
                <ClearIcon onClick={resetSearch} />
              </IconButton>
            ),
          }}
        />
        {isSmallScreen && (
          <>
            <SearchButton
              size={"small"}
              onClick={handleSearchButtonClick}
            />
            <SortButton
              availableSortCriteriaMappings={availableThreadSortCriteria}
              setSortCriteria={handleSortCriteriaClick}
              size={"small"}
            />
          </>
        )}
        {!isSmallScreen && (
          <>
            <SearchButton
              size={"large"}
              onClick={handleSearchButtonClick}
            />
            <SortButton
              availableSortCriteriaMappings={availableThreadSortCriteria}
              setSortCriteria={handleSortCriteriaClick}
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
      <div className={"mt-8"}>
        {isError && (
          <Alert
            severity={"error"}
            className={"mx-6 my-6"}
          >
            {"An error occurred while fetching threads " + errorMessage}
          </Alert>
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
      </div>
    </Box>
  );
}
