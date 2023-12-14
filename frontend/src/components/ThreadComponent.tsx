"use client";

import Thread from "@/models/thread/Thread";
import Box from "@mui/material/Box";
import { ListItemText } from "@mui/material";
import Link from "next/link";
import Typography from "@mui/material/Typography";
import UserAvatarDetails from "@/components/UserAvatarDetails";
import UserContentTimestamp from "@/components/UserContentTimestamp";
import ThreadTag from "@/components/ThreadTag";

// Creates a thread component to be displayed in the list
export default function ThreadComponent(t: Readonly<Thread>) {

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
              <UserAvatarDetails creator={t.creator} />
            </Box>
            <Box aria-label={"Thread Created Time"}
                 sx={{
                   display: "inline-flex",
                   gap: { xs: "0.4rem", sm: "0.5rem" },
                   alignItems: "center",
                 }}>
              <UserContentTimestamp
                createdTimestamp={t.created_time}
                updatedTimestamp={t.updated_time} />
            </Box>
          </Box>
          <Box aria-label={"Thread Tags and Comments"}
               sx={{
                 display: "flex",
                 flexWrap: "nowrap",
                 justifyContent: "space-between",
                 gap: "1rem",
                 maxWidth: "100%",
               }}>
            <Box aria-label={"Thread Tags"}
                 sx={{
                   display: "flex",
                   gap: "1rem",
                   overflow: "scroll",
                   mt: 2,
                 }}>
              {t.tags.map((tag) => (
                <ThreadTag key={t.id + tag} tag={tag} />
              ))}
            </Box>
            <Box aria-label={"Thread Comments"}
                 sx={{
                   display: "inline-flex",
                   gap: "1rem",
                   flexWrap: "nowrap",
                   mt: 2,
                   alignItems: "right",
                   flexShrink: 0,
                 }}>
              <Typography>{t.num_comments + " comments"}</Typography>
            </Box>
          </Box>
        </Box>
      </Link>
    </Box>
  );
}
