"use client";

import Thread from "@/models/thread/Thread";
import Box from "@mui/material/Box";
import { Chip, ListItemText } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import Link from "next/link";
import Typography from "@mui/material/Typography";
import { toSvg } from "jdenticon";

// Creates a thread component to be displayed in the list
export default function ThreadComponent(t: Readonly<Thread>) {

  const formattedCreatedTime = new Intl.DateTimeFormat("en-GB", {
    dateStyle: "short",
    timeStyle: "short",
    timeZone: "Asia/Singapore",
  }).format(t.created_time);

  function handleChipClick(tagName: string) {
    console.log(tagName);
  }

  return (
    <Box
      sx={{
        mx: { xs: 0, sm: 2 },
        my: 1,
        border: 2,
        borderColor: "lightgray",
        borderRadius: "0.375rem",
        width: "100%",
        overflow: "hidden",
      }}>
      <Link href={`/viewthread/${t.id}`}
            style={{ textDecoration: "none", color: "#000" }}>
        <Box sx={{ p: 2 }} aria-label={"Thread Display Box"}>
          <Box aria-label={"Thread Title"}
               sx={{ mb: 1 }}>
            <Typography sx={{
              fontSize: { xs: "1.25rem", sm: "1.5rem" },
              fontWeight: "600",
              textOverflow: "ellipsis",
              overflow: "hidden",
              display: "-webkit-box",
              WebkitLineClamp: "2",
              WebkitBoxOrient: "vertical",
            }}>{t.title}</Typography>
          </Box>
          <Box aria-label={"Thread Body"}
               sx={{ mb: 1 }}>
            <Typography sx={{
              fontSize: { xs: "1rem", sm: "1.25rem" },
              textOverflow: "ellipsis",
              overflow: "hidden",
              display: "-webkit-box",
              WebkitLineClamp: "2",
              WebkitBoxOrient: "vertical",
            }}>{t.body}</Typography>
          </Box>
          <Box aria-label={"Thread Details"}
               sx={{
                 display: "inline-flex",
                 flexDirection: { xs: "column", sm: "row" },
                 flexWrap: "wrap",
                 width: "100%",
                 mt: 1,
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
              <Avatar alt={t.creator}
                      src={`data:image/svg+xml;utf8,${encodeURIComponent(toSvg(t.creator, 50))}`}
                      sx={{
                        bgcolor: "white",
                        border: 1,
                        borderColor: "darkgray",
                        width: { xs: "20px", sm: "30px" },
                        height: { xs: "20px", sm: "30px" },
                        aspectRatio: 1,
                      }}>
              </Avatar>
              <ListItemText secondary={t.creator}
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
          <Box aria-label={"Thread Tags and Comments"}
               sx={{
                 display: "flex",
                 flexWrap: "nowrap",
                 justifyContent: "space-between",
                 gap: "1rem",
                 alignItems: "flex-end",
               }}>
            <Box aria-label={"Thread Tags"}
                 sx={{
                   display: "inline-flex",
                   gap: "1rem",
                   flexWrap: "wrap",
                   mt: 2,
                 }}>
              {t.tags.map((tag) => (
                <Chip key={t.id + tag} label={tag} onClick={() => handleChipClick(tag)} />
              ))}
            </Box>
            <Box aria-label={"Thread Comments"}
                 sx={{
                   display: "inline-flex",
                   gap: "1rem",
                   flexWrap: "wrap",
                   mt: 2,
                   alignItems: "right",
                 }}>
              <Typography>{t.num_comments + " comments"}</Typography>
            </Box>
          </Box>
        </Box>
      </Link>
    </Box>

  );
}
