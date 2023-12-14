"use client";

import Avatar from "@mui/material/Avatar";
import { toSvg } from "jdenticon";
import { ListItemText } from "@mui/material";

export default function UserAvatarDetails(props: Readonly<{
  creator: string,
  textColor?: string
  fontSize?: string
}>) {

  const textColor = props.textColor ?? "rgba(0, 0, 0, 0.6)";
  const fontSize = props.fontSize ?? "0.875rem";

  return (
    <>
      <Avatar alt={props.creator}
              src={`data:image/svg+xml;utf8,${encodeURIComponent(toSvg(props.creator, 50))}`}
              sx={{
                bgcolor: "white",
                border: 1,
                borderColor: "darkgray",
                width: { xs: "20px", sm: "30px" },
                height: { xs: "20px", sm: "30px" },
                aspectRatio: 1,
              }}>
      </Avatar>
      <ListItemText secondary={props.creator}
                    secondaryTypographyProps={{
                      sx: {
                        textOverflow: "ellipsis",
                        overflow: "hidden",
                        color: textColor,
                        fontSize: fontSize,
                      },
                    }}
      />
    </>);
}
