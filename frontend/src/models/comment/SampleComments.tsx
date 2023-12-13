import Comment from "@/models/comment/Comment";

export const SampleComments: Comment[] = [
  {
    id: "1",
    body: "wow a comment",
    creator: "user1",
    created_time: new Date("2021-10-17T15:00:00.000Z"),
    updated_time: new Date("2021-10-17T15:00:00.000Z"),
    thread_id: "1",
  },
  {
    id: "2",
    body: "Sample Comment 2",
    creator: "user2",
    created_time: new Date("2022-10-17T15:00:00.000Z"),
    updated_time: new Date("2022-10-17T15:00:00.000Z"),
    thread_id: "2",
  },
  {
    id: "3",
    body: "Sample Content 3",
    creator: "user3",
    created_time: new Date("2023-10-17T15:00:00.000Z"),
    updated_time: new Date("2023-10-17T15:00:00.000Z"),
    thread_id: "3",
  },
  {
    id: "4",
    body: "\"Many of the case competitions we’ve joined in our college career were solely focused on business\n" +
      " profits. This challenge was especially meaningful to us because of its mission. We were honored to be a part of something that could ultimately help millions of people from all parts of the world.\”",
    creator: "wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww",
    created_time: new Date("2023-10-17T15:00:00.000Z"),
    updated_time: new Date("2023-10-17T15:00:00.000Z"),
    thread_id: "4",
  },
];

export default SampleComments;
