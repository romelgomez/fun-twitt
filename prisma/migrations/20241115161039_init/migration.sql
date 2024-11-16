-- CreateTable
CREATE TABLE "User" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "sort_id" TEXT NOT NULL,
    "name" TEXT NOT NULL DEFAULT '',
    "created" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated" DATETIME NOT NULL,
    "deleted" DATETIME
);

-- CreateIndex
CREATE UNIQUE INDEX "User_sort_id_key" ON "User"("sort_id");
