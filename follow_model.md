# Follow model

The `FollowerID` and `FolloweeID` fields represent the relationship between two users in the `Follow` model. Here’s a detailed description of each:

---

### **FollowerID**
- **Definition:** 
  - This field holds the `id` of the user who is following another user.
  
- **Purpose:**
  - Identifies the user initiating the follow action.
  - Represents the "follower" in the relationship.

- **Example:**
  - If User A follows User B:
    - `FollowerID` will contain the `id` of User A.

---

### **FolloweeID**
- **Definition:** 
  - This field holds the `id` of the user being followed.
  
- **Purpose:**
  - Identifies the user being followed by another user.
  - Represents the "followee" in the relationship.

- **Example:**
  - If User A follows User B:
    - `FolloweeID` will contain the `id` of User B.

---

### **Illustration of the Relationship**

Imagine the following scenario:
- User A (`id: user_ABC123`) follows User B (`id: user_DEF456`).

In the `Follow` table:
```json
{
  "id": "follow_XYZ789",
  "follower_id": "user_ABC123",
  "followee_id": "user_DEF456",
  "created_at": "2024-11-15T12:00:00Z"
}
```

---

### **Usage in Queries:**
1. **Retrieve All Followers of a User (Followees):**
   - Query all `Follow` records where `FolloweeID` equals the target user's `id`.
   - Example: Find who is following User B.
     ```sql
     SELECT * FROM Follow WHERE followee_id = 'user_DEF456';
     ```

2. **Retrieve All Followees of a User (Followers):**
   - Query all `Follow` records where `FollowerID` equals the target user's `id`.
   - Example: Find whom User A is following.
     ```sql
     SELECT * FROM Follow WHERE follower_id = 'user_ABC123';
     ```

---

### Summary:
- **`FollowerID`:** The user who follows someone.
- **`FolloweeID`:** The user being followed.
