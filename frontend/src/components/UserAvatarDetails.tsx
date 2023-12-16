import { toSvg } from "jdenticon";
import Typography from "@mui/material/Typography";

export default function UserAvatarDetails(
  props: Readonly<{
    creator: string;
    textColor?: string;
    fontSize?: string;
  }>,
) {
  // TODO: Fix this
  // const textColor = props.textColor ?? "rgba(0, 0, 0, 0.6)";
  // const fontSize = props.fontSize ?? "0.875rem";

  return (
    <div className={"flex w-full items-center"}>
      <img
        alt={props.creator}
        src={`data:image/svg+xml;utf8,${encodeURIComponent(toSvg(props.creator, 50))}`}
        className={"border-1 mr-2 flex aspect-square h-7 w-7 rounded-full border-gray-500 bg-white"}
      />
      <Typography
        noWrap
        className={"overflow-hidden text-ellipsis"}
      >
        {props.creator}
      </Typography>
    </div>
  );
}
