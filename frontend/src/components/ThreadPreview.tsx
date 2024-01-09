import Box from "@mui/material/Box";
import { ListItemText } from "@mui/material";
import Typography from "@mui/material/Typography";
import Thread from "../models/thread/Thread.tsx";
import { Link } from "react-router-dom";
import UserAvatarDetails from "./UserAvatarDetails.tsx";
import UserContentTimestamp from "./UserContentTimestamp.tsx";
import ThreadTag from "./ThreadTag.tsx";

// Creates a thread component to be displayed in the list
export default function ThreadPreview(t: Readonly<Thread>) {
  return (
    <Box className={"xs:mx-0 mx-2 my-2 w-full overflow-hidden rounded-md border-2 border-solid border-gray-300"}>
      <Link
        to={`/viewthread/${t.id}`}
        style={{ textDecoration: "none", color: "#000" }}
      >
        <Box
          sx={{ p: 2 }}
          aria-label={"Thread Display Box"}
        >
          <Box
            aria-label={"Thread Title"}
            className={"mb-2"}
          >
            <Typography
              sx={{
                fontSize: { xs: "1.125rem", sm: "1.25rem" },
                fontWeight: "600",
                textOverflow: "ellipsis",
                overflow: "hidden",
                display: "-webkit-box",
                WebkitLineClamp: "2",
                WebkitBoxOrient: "vertical",
              }}
            >
              {t.title}
            </Typography>
          </Box>
          <Box
            aria-label={"Thread Body"}
            className={"mb-4"}
          >
            <Typography
              sx={{
                fontSize: { xs: "0.875rem", sm: "1rem" },
                textOverflow: "ellipsis",
                overflow: "hidden",
                display: "-webkit-box",
                WebkitLineClamp: "2",
                WebkitBoxOrient: "vertical",
              }}
            >
              {t.body}
            </Typography>
          </Box>
          <Box
            aria-label={"Thread Details"}
            className={"xs:flex-col flex flex-row flex-wrap align-middle"}
          >
            <Box
              aria-label={"Thread Creator"}
              sx={{
                display: "flex",
                gap: { xs: "0.4rem", sm: "0.5rem" },
                alignItems: "center",
                pr: 1,
                overflow: "hidden",
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
              <UserAvatarDetails creator={t.creator} />
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
                createdTimestamp={new Date(t.created_time)}
                updatedTimestamp={new Date(t.updated_time)}
              />
            </Box>
          </Box>
          <Box
            aria-label={"Thread Tags and Comments"}
            sx={{
              display: "flex",
              flexWrap: "nowrap",
              justifyContent: "space-between",
              gap: "1rem",
              maxWidth: "100%",
              alignItems: "center",
            }}
          >
            <Box
              aria-label={"Thread Tags"}
              className={"mt-4 flex gap-3 overflow-scroll"}
            >
              {t.tags.map((tag) => (
                <ThreadTag
                  key={t.id + tag}
                  tag={tag}
                />
              ))}
            </Box>
            <Box
              aria-label={"Thread Comments"}
              sx={{
                display: "inline-flex",
                gap: "1rem",
                flexWrap: "nowrap",
                mt: 2,
                alignItems: "right",
                flexShrink: 0,
              }}
            >
              <Typography>{t.num_comments + " comments"}</Typography>
            </Box>
          </Box>
        </Box>
      </Link>
    </Box>
  );
}
