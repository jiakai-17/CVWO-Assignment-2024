import { ListItemText, Tooltip } from "@mui/material";

// Formats and displays the timestamp of a user's content
export default function UserContentTimestamp(
  props: Readonly<{
    createdTimestamp: Date;
    updatedTimestamp: Date;
  }>,
) {
  const dateFormatOptions: Intl.DateTimeFormatOptions = {
    year: "numeric",
    month: "numeric",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
    hour12: false,
  };

  const formattedCreatedTime = props.createdTimestamp.toLocaleString("en-SG", dateFormatOptions);

  let displayString = ` on ${formattedCreatedTime}`;
  let tooltipString = `Created on ${props.createdTimestamp.toString()}`;

  if (props.updatedTimestamp !== undefined && props.updatedTimestamp.getTime() !== props.createdTimestamp.getTime()) {
    displayString += ` (edited)`;
    tooltipString += `\n\n Edited on ${props.updatedTimestamp.toString()}`;
  }

  return (
    <Tooltip title={<p style={{ whiteSpace: "pre-line", margin: 0 }}>{tooltipString}</p>}>
      <ListItemText
        secondary={displayString}
        sx={{ fontSize: { xs: "0.5rem", sm: "body2" } }}
      />
    </Tooltip>
  );
}
