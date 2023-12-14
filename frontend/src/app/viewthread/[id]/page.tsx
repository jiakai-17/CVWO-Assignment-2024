"use client";

import sampleThreads from "@/models/thread/SampleThreads";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Link from "next/link";
import { Divider, ListItemText } from "@mui/material";
import TextField from "@mui/material/TextField";
import * as React from "react";
import { useEffect, useState } from "react";
import sampleComments from "@/models/comment/SampleComments";
import UserAvatarDetails from "@/components/UserAvatarDetails";
import UserContentTimestamp from "@/components/UserContentTimestamp";
import ThreadTag from "@/components/ThreadTag";
import SortButton from "@/components/SortButton";
import CommentTextField from "@/components/CommentTextField";

export default function Page({ params }: Readonly<{ params: { id: string } }>) {

  const threadToDisplay = sampleThreads.find(thread => thread.id === params.id);

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
    <Box>
      {threadToDisplay === undefined &&
        <Box sx={{
          mx: 2,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}>
          <Typography variant="h4" sx={{ mb: 6 }}>Thread Not Found</Typography>
          <Link href={"/"}>
            <Button
              variant="contained"
              disableElevation
              size="large">
              Back to Home
            </Button>
          </Link>
        </Box>
      }

      {threadToDisplay !== undefined &&
        <>
          <Box sx={{ mx: 4, mb: 4 }}>
            <Link href={"/"}>
              <Button
                variant="outlined"
                disableElevation
                size="large">
                Back to Home
              </Button>
            </Link>
          </Box>
          <Box sx={{
            mx: 4,
            maxWidth: "full",
            borderRadius: "0.375rem",
            border: 2,
            borderColor: "lightgray",
            px: 3,
          }}>
            <Typography sx={{
              fontSize: { xs: "1.25rem", sm: "1.5rem" },
              fontWeight: "600",
              py: 2,
              overflowWrap: "break-word",
            }}>{threadToDisplay.title}</Typography>
            <Divider />
            <Typography sx={{
              fontSize: { xs: "1rem", sm: "1.25rem" },
              overflowWrap: "break-word",
              py: 2,
              mb: 5,
            }}>{threadToDisplay.body}</Typography>
            <Divider />
            <Box sx={{
              display: "inline-flex",
              gap: "0.5rem",
              flexWrap: "wrap",
              alignItems: "center",
            }}>
              <Typography sx={{
                fontSize: "1rem",
                py: 2,
                color: "rgba(0, 0, 0, 0.6)",
              }}>Tags:
              </Typography>
              {threadToDisplay.tags.map(tag => (
                <ThreadTag key={threadToDisplay.id + tag} tag={tag} />
              ))}
            </Box>
            <Divider />
            <Box aria-label={"Thread Details"}
                 sx={{
                   display: "inline-flex",
                   flexDirection: { xs: "column", sm: "row" },
                   flexWrap: "wrap",
                   width: "100%",
                   my: 2,
                 }}>
              <Box aria-label={"Thread Creator"}
                   sx={{
                     display: "flex",
                     gap: { xs: "0.4rem", sm: "0.5rem" },
                     alignItems: "center",
                     maxWidth: "100%",
                     pr: 1,
                   }}>
                <ListItemText secondary={"Posted by "}
                              sx={{
                                fontSize: { xs: "0.5rem", sm: "body2" },
                                flexShrink: 0,
                                flexGrow: 0,
                              }} />
                <UserAvatarDetails creator={threadToDisplay.creator} />
              </Box>
              <Box aria-label={"Thread Created Time"}
                   sx={{
                     display: "inline-flex",
                     gap: { xs: "0.4rem", sm: "0.5rem" },
                     alignItems: "center",
                   }}>
                <UserContentTimestamp
                  createdTimestamp={threadToDisplay.created_time}
                  updatedTimestamp={threadToDisplay.updated_time} />
              </Box>
            </Box>
          </Box>
          <Divider sx={{ mx: 4, my: 4 }} />
          <CommentTextField
            setCommentContent={setCommentContent}
            handleSubmit={() => console.log("submit Comment", commentContent)} />
          <Divider sx={{ mx: 4, mt: 2, mb: 4 }} />
          <Box
            sx={{ mx: 4, display: "flex", justifyContent: "space-between", alignItems: "center" }}>
            <Typography variant="h5">99 Comments</Typography>
            <SortButton
              availableSortCriteriaMappings={availableCommentSortCriteria}
              setSortCriteria={setCommentSortCriteria}
              size={"large"} />
          </Box>
          <Box sx={{ mx: 4, mb: 20 }}>
            {sampleComments.map(comment => (
              <Box key={comment.id}
                   sx={{
                     my: 4,
                     p: 2,
                     border: 2,
                     borderColor: "lightgray",
                     borderRadius: "0.375rem",
                   }}>
                <Box sx={{
                  display: "flex",
                  justifyContent: "space-between",
                  mb: 2,
                  maxWidth: "100%",
                }}>
                  <Box sx={{
                    display: "flex",
                    gap: { xs: "0.4rem", sm: "0.5rem" },
                    alignItems: "center",
                    pr: 1,
                    minWidth: 0,
                  }}>
                    <UserAvatarDetails creator={comment.creator} />
                  </Box>
                  <Box aria-label={"Thread Created Time"}
                       sx={{
                         display: "inline-flex",
                         gap: { xs: "0.4rem", sm: "0.5rem" },
                         alignItems: "center",
                         flexShrink: 0,
                       }}>
                    <UserContentTimestamp
                      createdTimestamp={comment.created_time}
                      updatedTimestamp={comment.updated_time}
                    />
                  </Box>
                </Box>
                <Typography>
                  {comment.body}
                </Typography>
              </Box>
            ))}
          </Box>
        </>
      }
    </Box>
  );
}
