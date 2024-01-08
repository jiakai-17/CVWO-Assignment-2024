import Thread from "@/models/thread/Thread";

export const SampleThreads: Thread[] = [
  {
    id: "1",
    title: "Sample ThreadComponent 1",
    body: "Sample Content 1",
    creator: "user1",
    created_time: new Date("2021-10-17T15:00:00.000Z"),
    updated_time: new Date("2021-10-18T15:00:00.000Z"),
    num_comments: 0,
    tags: ["tag1", "tag2"],
  },
  {
    id: "2",
    title: "Sample ThreadComponent 2",
    body: "Sample Content 2",
    creator: "user2",
    created_time: new Date("2022-10-17T15:00:00.000Z"),
    updated_time: new Date("2022-10-17T15:00:00.000Z"),
    num_comments: 1000,
    tags: ["tag1", "tag3"],
  },
  {
    id: "3",
    title: "Sample ThreadComponent 3",
    body: "Sample Content 3",
    creator: "user3",
    created_time: new Date("2023-10-17T15:00:00.000Z"),
    updated_time: new Date("2023-10-17T15:00:00.000Z"),
    num_comments: 26,
    tags: ["tag2", "tag3"],
  },
  {
    id: "4",
    title: "Sample wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww 3",
    body: "\"Many of the case competitions we’ve joined in our college career were solely focused on business\n" +
      " profits. This challenge was especially meaningful to us because of its mission. We were honored to be a part of something that could ultimately help millions of people from all parts of the world.\”",
    creator: "wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww",
    created_time: new Date("2023-10-17T15:00:00.000Z"),
    updated_time: new Date("2023-10-17T15:00:00.000Z"),
    num_comments: 26,
    tags: ["tag2", "tag3", "tag3", "tag3", "tag3", "tag3", "tag3", "tag3", "tag3", "tag3", "tag3", "tag3", "tag3"],
  },
];

export default SampleThreads;
