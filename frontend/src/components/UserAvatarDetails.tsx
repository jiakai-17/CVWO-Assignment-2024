import { toSvg } from "jdenticon";
import Typography from "@mui/material/Typography";

// Renders a user avatar with the username
export default function UserAvatarDetails(
  props: Readonly<{
    creator: string;
    textColor?: string;
    fontSize?: string;
  }>,
) {
  return (
    <div className={"flex items-center overflow-hidden text-ellipsis"}>
      <img
        alt={props.creator}
        src={`data:image/svg+xml;utf8,${encodeURIComponent(toSvg(props.creator, 50))}`}
        className={"mr-2 flex aspect-square h-7 w-7 rounded-full border border-solid border-gray-500 bg-white"}
      />
      <Typography
        noWrap
        className={"overflow-hidden text-ellipsis"}
        style={{
          color: props.textColor ?? "rgba(0, 0, 0, 0.6)",
          fontSize: props.fontSize ?? "0.875rem",
        }}
      >
        {props.creator}
      </Typography>
    </div>
  );
}
