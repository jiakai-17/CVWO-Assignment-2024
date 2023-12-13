"use client";

import sampleThreads from "@/models/thread/SampleThreads";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Link from "next/link";
import { Chip, Divider, ListItemText } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import { toSvg } from "jdenticon";
import TextField from "@mui/material/TextField";
import * as React from "react";
import SortIcon from "@mui/icons-material/Sort";
import sampleComments from "@/models/comment/SampleComments";

export default function Page({ params }: Readonly<{ params: { id: string } }>) {

  const threadToDisplay = sampleThreads.find(thread => thread.id === params.id);


  const formattedCreatedTime = threadToDisplay === undefined ? "" :
    new Intl.DateTimeFormat("en-GB", {
      dateStyle: "short",
      timeStyle: "short",
      timeZone: "Asia/Singapore",
    }).format(threadToDisplay.created_time);

  function handleChipClick(tagName: string) {
    console.log(tagName);
  }

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
                <Chip key={threadToDisplay.id + tag} label={tag}
                      onClick={() => handleChipClick(tag)} />
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
                <Avatar alt={threadToDisplay.creator}
                        src={`data:image/svg+xml;utf8,${encodeURIComponent(toSvg(threadToDisplay.creator, 50))}`}
                        sx={{
                          bgcolor: "white",
                          border: 1,
                          borderColor: "darkgray",
                          width: { xs: "20px", sm: "30px" },
                          height: { xs: "20px", sm: "30px" },
                          aspectRatio: 1,
                        }}>
                </Avatar>
                <ListItemText secondary={threadToDisplay.creator}
                              secondaryTypographyProps={{
                                sx: {
                                  textOverflow: "ellipsis",
                                  overflow: "hidden",
                                },
                              }}
                />
              </Box>
              <Box aria-label={"Thread Created Time"}
                   sx={{
                     display: "inline-flex",
                     gap: { xs: "0.4rem", sm: "0.5rem" },
                     alignItems: "center",
                   }}>
                <div>
                  <ListItemText secondary={" on " + formattedCreatedTime}
                                sx={{ fontSize: { xs: "0.5rem", sm: "body2" } }} />
                </div>
              </Box>
            </Box>
          </Box>
          <Divider sx={{ mx: 4, my: 4 }} />
          <Box sx={{ mx: 4, width: "full" }}>
            <TextField
              id="outlined-multiline-static"
              label="Leave a comment..."
              multiline
              rows={3}
              variant="outlined"
              fullWidth
            />
            <Box sx={{ display: "flex", justifyContent: "flex-end" }}>
              <Button
                variant="contained"
                size={"large"}
                sx={{ my: 2 }}
              >
                COMMENT
              </Button>
            </Box>
          </Box>
          <Divider sx={{ mx: 4, mt: 2, mb: 4 }} />
          <Box
            sx={{ mx: 4, display: "flex", justifyContent: "space-between", alignItems: "center" }}>
            <Typography variant="h5">99 Comments</Typography>
            <Button
              variant="outlined"
              size="large"
              endIcon={<SortIcon />}>
              Sort...
            </Button>
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
                     overflow: "hidden",
                   }}>
                <Box sx={{ display: "flex", justifyContent: "space-between", mb: 2 }}>
                  <Box sx={{
                    display: "flex",
                    gap: { xs: "0.4rem", sm: "0.5rem" },
                    alignItems: "center",
                    maxWidth: "100%",
                    pr: 1,
                  }}>
                    <Avatar alt={comment.creator}
                            src={`data:image/svg+xml;utf8,${encodeURIComponent(toSvg(comment.creator, 50))}`}
                            sx={{
                              bgcolor: "white",
                              border: 1,
                              borderColor: "darkgray",
                              width: { xs: "20px", sm: "30px" },
                              height: { xs: "20px", sm: "30px" },
                              aspectRatio: 1,
                            }}>
                    </Avatar>
                    <ListItemText secondary={comment.creator}
                                  secondaryTypographyProps={{
                                    sx: {
                                      textOverflow: "ellipsis",
                                      overflow: "hidden",
                                    },
                                  }}
                    />
                  </Box>
                  <Box aria-label={"Thread Created Time"}
                       sx={{
                         display: "inline-flex",
                         gap: { xs: "0.4rem", sm: "0.5rem" },
                         alignItems: "center",
                       }}>
                    <div>
                      <ListItemText secondary={" on " + formattedCreatedTime}
                                    sx={{ fontSize: { xs: "0.5rem", sm: "body2" } }} />
                    </div>
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
