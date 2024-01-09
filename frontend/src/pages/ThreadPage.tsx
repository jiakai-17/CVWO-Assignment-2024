import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import { Divider, ListItemText } from "@mui/material";
import { useEffect, useState } from "react";
import { Link, Params, useLoaderData } from "react-router-dom";
import sampleThreads from "../models/thread/SampleThreads.tsx";
import ThreadTag from "../components/ThreadTag.tsx";
import UserContentTimestamp from "../components/UserContentTimestamp.tsx";
import UserAvatarDetails from "../components/UserAvatarDetails.tsx";
import SortButton from "../components/SortButton.tsx";
import sampleComments from "../models/comment/SampleComments.tsx";
import Thread from "../models/thread/Thread.tsx";
import CommentTextField from "../components/CommentTextField.tsx";
import { ThreadComment } from "./ThreadComment.tsx";

export async function loader({ params }: { params: Params<"id"> }) {
  return sampleThreads.find((thread) => thread.id === params.id) ?? null;
}

export default function ThreadPage() {
  const threadToDisplay = useLoaderData() as Thread;

  const [commentContent, setCommentContent] = useState("");

  useEffect(() => {
    console.log("comment content changed", commentContent);
  }, [commentContent]);

  // Sorting comments by criteria
  const [commentSortCriteria, setCommentSortCriteria] = useState("Newest first");

  const availableCommentSortCriteria = new Map<string, string>([
    ["Newest first", "date desc"],
    ["Oldest first", "date asc"],
  ]);

  useEffect(() => {
    console.log(commentSortCriteria);
  }, [commentSortCriteria]);

  return (
    <Box className={"mt-16"}>
      {threadToDisplay === null && (
        <Box
          sx={{
            mx: 2,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Typography
            variant="h4"
            sx={{ mb: 6 }}
          >
            Thread Not Found
          </Typography>
          <Link to={"/"}>
            <Button
              variant="contained"
              disableElevation
              size="large"
            >
              Back to Home
            </Button>
          </Link>
        </Box>
      )}

      {threadToDisplay !== null && (
        <>
          <Box sx={{ mx: 4, mb: 4 }}>
            <Link to={"/"}>
              <Button
                variant="outlined"
                disableElevation
                size="large"
              >
                Back to Home
              </Button>
            </Link>
          </Box>
          <Box
            sx={{
              mx: 4,
              maxWidth: "full",
              borderRadius: "0.375rem",
              border: 2,
              borderColor: "lightgray",
              px: 3,
            }}
          >
            <Typography
              sx={{
                fontSize: { xs: "1.125rem", sm: "1.25rem" },
                fontWeight: "600",
                py: 2,
                overflowWrap: "break-word",
              }}
            >
              {threadToDisplay.title}
            </Typography>
            <Divider />
            <Typography
              sx={{
                fontSize: { xs: "0.875rem", sm: "1rem" },
                overflowWrap: "break-word",
                py: 2,
                mb: 5,
              }}
            >
              {threadToDisplay.body}
            </Typography>
            <Divider />
            <Box
              sx={{
                display: "inline-flex",
                gap: "0.5rem",
                flexWrap: "wrap",
                alignItems: "center",
              }}
            >
              <Typography
                sx={{
                  fontSize: "1rem",
                  py: 2,
                  color: "rgba(0, 0, 0, 0.6)",
                }}
              >
                Tags:
              </Typography>
              {threadToDisplay.tags.map((tag) => (
                <ThreadTag
                  key={threadToDisplay.id + tag}
                  tag={tag}
                />
              ))}
            </Box>
            <Divider />
            <Box
              aria-label={"Thread Details"}
              sx={{
                display: "inline-flex",
                flexDirection: { xs: "column", sm: "row" },
                flexWrap: "wrap",
                width: "100%",
                my: 2,
              }}
            >
              <Box
                aria-label={"Thread Creator"}
                sx={{
                  display: "flex",
                  gap: { xs: "0.4rem", sm: "0.5rem" },
                  alignItems: "center",
                  maxWidth: "100%",
                  pr: 1,
                }}
              >
                <ListItemText
                  secondary={"Posted by "}
                  sx={{
                    fontSize: { xs: "0.5rem", sm: "body2" },
                    flexShrink: 0,
                    flexGrow: 0,
                  }}
                />
                <UserAvatarDetails creator={threadToDisplay.creator} />
              </Box>
              <Box
                aria-label={"Thread Created Time"}
                sx={{
                  display: "inline-flex",
                  gap: { xs: "0.4rem", sm: "0.5rem" },
                  alignItems: "center",
                }}
              >
                <UserContentTimestamp
                  createdTimestamp={threadToDisplay.created_time}
                  updatedTimestamp={threadToDisplay.updated_time}
                />
              </Box>
            </Box>
            {threadToDisplay.creator === localStorage.getItem("username") && (
              <>
                <Divider />
                <Box className={"my-4 flex items-center justify-end"}>
                  <Link to={`/viewthread/${threadToDisplay.id}/edit`}>
                    <Button
                      variant="outlined"
                      size="large"
                      sx={{ mr: 2 }}
                    >
                      Edit
                    </Button>
                  </Link>
                  <Button
                    variant="outlined"
                    size="large"
                    color="error"
                    sx={{ mr: 2 }}
                  >
                    Delete
                  </Button>
                </Box>
              </>
            )}
          </Box>

          <Divider sx={{ mx: 4, my: 4 }} />
          <CommentTextField
            setCommentContent={setCommentContent}
            handleSubmit={() => console.log("submit Comment", commentContent)}
          />
          <Divider sx={{ mx: 4, mt: 2, mb: 4 }} />
          <Box sx={{ mx: 4, display: "flex", justifyContent: "space-between", alignItems: "center" }}>
            <Typography variant="h5">99 Comments</Typography>
            <SortButton
              availableSortCriteriaMappings={availableCommentSortCriteria}
              setSortCriteria={setCommentSortCriteria}
              size={"large"}
            />
          </Box>
          <Box sx={{ mx: 4, mb: 20 }}>
            {sampleComments.map((comment) => (
              <ThreadComment
                key={comment.id}
                comment={comment}
              />
            ))}
          </Box>
        </>
      )}
    </Box>
  );
}
