#!/bin/bash

BASE_URL="http://localhost:8080/api/v1"

echo "============================================"
echo "Testing Kurikulum Backend API Endpoints"
echo "============================================"
echo ""

# Test Health
echo "1. Testing Health Check..."
HEALTH=$(curl -s "$BASE_URL/health")
if echo "$HEALTH" | grep -q '"status":"ok"'; then
    echo "   ✓ Health check passed"
else
    echo "   ✗ Health check failed"
fi

# Register a test kaprodi user for full testing
echo ""
echo "2. Testing Register (Kaprodi)..."
REGISTER=$(curl -s -X POST "$BASE_URL/auth/register" \
    -H "Content-Type: application/json" \
    -d '{"nama":"Kaprodi Test","email":"kaprodi@test.com","password":"Test1234!","role":"kaprodi"}' 2>/dev/null)
if echo "$REGISTER" | grep -q '"success":true'; then
    echo "   ✓ Register passed"
elif echo "$REGISTER" | grep -q 'sudah terdaftar'; then
    echo "   ✓ Register skipped (user exists)"
else
    echo "   ? Register: $REGISTER"
fi

# Login as kaprodi
echo ""
echo "3. Testing Login..."
LOGIN=$(curl -s -X POST "$BASE_URL/auth/login" \
    -H "Content-Type: application/json" \
    -d '{"email":"kaprodi@test.com","password":"Test1234!"}')

if echo "$LOGIN" | grep -q '"access_token"'; then
    TOKEN=$(echo "$LOGIN" | grep -o '"access_token":"[^"]*"' | cut -d'"' -f4)
    REFRESH_TOKEN=$(echo "$LOGIN" | grep -o '"refresh_token":"[^"]*"' | cut -d'"' -f4)
    echo "   ✓ Login passed"
else
    echo "   ✗ Login failed: $LOGIN"
    exit 1
fi

AUTH="Authorization: Bearer $TOKEN"

# Test Profile
echo ""
echo "4. Testing Get Profile..."
PROFILE=$(curl -s "$BASE_URL/auth/profile" -H "$AUTH")
if echo "$PROFILE" | grep -q '"success":true'; then
    echo "   ✓ Get Profile passed"
else
    echo "   ✗ Get Profile failed: $PROFILE"
fi

# Test Update Profile
echo ""
echo "5. Testing Update Profile..."
UPDATE_PROFILE=$(curl -s -X PUT "$BASE_URL/auth/profile" \
    -H "$AUTH" \
    -H "Content-Type: application/json" \
    -d '{"nama":"Kaprodi Updated"}')
if echo "$UPDATE_PROFILE" | grep -q '"success":true'; then
    echo "   ✓ Update Profile passed"
else
    echo "   ✗ Update Profile failed: $UPDATE_PROFILE"
fi

# Test Dashboard
echo ""
echo "6. Testing Dashboard..."
DASHBOARD=$(curl -s "$BASE_URL/dashboard" -H "$AUTH")
if echo "$DASHBOARD" | grep -q '"success":true'; then
    echo "   ✓ Dashboard passed"
else
    echo "   ✗ Dashboard failed: $DASHBOARD"
fi

# Test Kaprodi Dashboard
echo ""
echo "7. Testing Kaprodi Dashboard..."
KAPRODI_DASH=$(curl -s "$BASE_URL/dashboard/kaprodi" -H "$AUTH")
if echo "$KAPRODI_DASH" | grep -q '"success":true'; then
    echo "   ✓ Kaprodi Dashboard passed"
else
    echo "   ✗ Kaprodi Dashboard failed: $KAPRODI_DASH"
fi

# Test Users
echo ""
echo "8. Testing Get All Users..."
USERS=$(curl -s "$BASE_URL/users" -H "$AUTH")
if echo "$USERS" | grep -q '"success":true'; then
    echo "   ✓ Get All Users passed"
else
    echo "   ✗ Get All Users failed: $USERS"
fi

# Test Get Dosen
echo ""
echo "9. Testing Get Dosen..."
DOSEN=$(curl -s "$BASE_URL/users/dosen" -H "$AUTH")
if echo "$DOSEN" | grep -q '"success":true'; then
    echo "   ✓ Get Dosen passed"
else
    echo "   ✗ Get Dosen failed: $DOSEN"
fi

# Test CPL
echo ""
echo "10. Testing CPL endpoints..."

# Create CPL
CPL_CREATE=$(curl -s -X POST "$BASE_URL/cpl" \
    -H "$AUTH" \
    -H "Content-Type: application/json" \
    -d '{"kode":"CPL-TEST-001","nama":"Capaian Pembelajaran Test","deskripsi":"Deskripsi capaian pembelajaran lulusan pertama"}')
if echo "$CPL_CREATE" | grep -q '"success":true'; then
    CPL_ID=$(echo "$CPL_CREATE" | grep -o '"id":"[^"]*"' | head -1 | cut -d'"' -f4)
    echo "   ✓ Create CPL passed (ID: $CPL_ID)"
elif echo "$CPL_CREATE" | grep -q 'sudah ada\|sudah digunakan\|duplicate'; then
    echo "   ✓ Create CPL skipped (already exists)"
else
    echo "   ✗ Create CPL failed: $CPL_CREATE"
fi

# Get All CPL
CPL_LIST=$(curl -s "$BASE_URL/cpl" -H "$AUTH")
if echo "$CPL_LIST" | grep -q '"success":true'; then
    echo "   ✓ Get All CPL passed"
else
    echo "   ✗ Get All CPL failed: $CPL_LIST"
fi

# Get CPL Statistics
CPL_STATS=$(curl -s "$BASE_URL/cpl/statistics" -H "$AUTH")
if echo "$CPL_STATS" | grep -q '"success":true'; then
    echo "   ✓ Get CPL Statistics passed"
else
    echo "   ✗ Get CPL Statistics failed: $CPL_STATS"
fi

# Get Active CPL
CPL_ACTIVE=$(curl -s "$BASE_URL/cpl/active" -H "$AUTH")
if echo "$CPL_ACTIVE" | grep -q '"success":true'; then
    echo "   ✓ Get Active CPL passed"
else
    echo "   ✗ Get Active CPL failed: $CPL_ACTIVE"
fi

# Test Mata Kuliah
echo ""
echo "11. Testing Mata Kuliah endpoints..."

# Create Mata Kuliah
MK_CREATE=$(curl -s -X POST "$BASE_URL/mata-kuliah" \
    -H "$AUTH" \
    -H "Content-Type: application/json" \
    -d '{"kode":"MK-TEST-001","nama":"Pemrograman Web Test","sks":3,"semester":3,"is_active":true}')
if echo "$MK_CREATE" | grep -q '"success":true'; then
    MK_ID=$(echo "$MK_CREATE" | grep -o '"id":"[^"]*"' | head -1 | cut -d'"' -f4)
    echo "   ✓ Create Mata Kuliah passed (ID: $MK_ID)"
elif echo "$MK_CREATE" | grep -q 'sudah ada\|sudah digunakan\|duplicate'; then
    echo "   ✓ Create Mata Kuliah skipped (already exists)"
else
    echo "   ✗ Create Mata Kuliah failed: $MK_CREATE"
fi

# Get All Mata Kuliah
MK_LIST=$(curl -s "$BASE_URL/mata-kuliah" -H "$AUTH")
if echo "$MK_LIST" | grep -q '"success":true'; then
    echo "   ✓ Get All Mata Kuliah passed"
else
    echo "   ✗ Get All Mata Kuliah failed: $MK_LIST"
fi

# Get My Mata Kuliah
MK_MY=$(curl -s "$BASE_URL/mata-kuliah/my" -H "$AUTH")
if echo "$MK_MY" | grep -q '"success":true'; then
    echo "   ✓ Get My Mata Kuliah passed"
else
    echo "   ✗ Get My Mata Kuliah failed: $MK_MY"
fi

# Test Notifications
echo ""
echo "12. Testing Notification endpoints..."

NOTIF_LIST=$(curl -s "$BASE_URL/notifications" -H "$AUTH")
if echo "$NOTIF_LIST" | grep -q '"success":true'; then
    echo "   ✓ Get Notifications passed"
else
    echo "   ✗ Get Notifications failed: $NOTIF_LIST"
fi

NOTIF_UNREAD=$(curl -s "$BASE_URL/notifications/unread-count" -H "$AUTH")
if echo "$NOTIF_UNREAD" | grep -q '"success":true'; then
    echo "   ✓ Get Unread Count passed"
else
    echo "   ✗ Get Unread Count failed: $NOTIF_UNREAD"
fi

# Test Documents
echo ""
echo "13. Testing Document endpoints..."

DOC_LIST=$(curl -s "$BASE_URL/documents" -H "$AUTH")
if echo "$DOC_LIST" | grep -q '"success":true'; then
    echo "   ✓ Get Documents passed"
else
    echo "   ✗ Get Documents failed: $DOC_LIST"
fi

DOC_TEMPLATES=$(curl -s "$BASE_URL/documents/templates" -H "$AUTH")
if echo "$DOC_TEMPLATES" | grep -q '"success":true'; then
    echo "   ✓ Get Document Templates passed"
else
    echo "   ✗ Get Document Templates failed: $DOC_TEMPLATES"
fi

# Test RPS
echo ""
echo "14. Testing RPS endpoints..."

RPS_LIST=$(curl -s "$BASE_URL/rps" -H "$AUTH")
if echo "$RPS_LIST" | grep -q '"success":true'; then
    echo "   ✓ Get All RPS passed"
else
    echo "   ✗ Get All RPS failed: $RPS_LIST"
fi

RPS_MY=$(curl -s "$BASE_URL/rps/my" -H "$AUTH")
if echo "$RPS_MY" | grep -q '"success":true'; then
    echo "   ✓ Get My RPS passed"
else
    echo "   ✗ Get My RPS failed: $RPS_MY"
fi

# Test CPL Assignments
echo ""
echo "15. Testing CPL Assignment endpoints..."

ASSIGN_LIST=$(curl -s "$BASE_URL/cpl-assignments" -H "$AUTH")
if echo "$ASSIGN_LIST" | grep -q '"success":true'; then
    echo "   ✓ Get All Assignments passed"
else
    echo "   ✗ Get All Assignments failed: $ASSIGN_LIST"
fi

ASSIGN_MY=$(curl -s "$BASE_URL/cpl-assignments/my" -H "$AUTH")
if echo "$ASSIGN_MY" | grep -q '"success":true'; then
    echo "   ✓ Get My Assignments passed"
else
    echo "   ✗ Get My Assignments failed: $ASSIGN_MY"
fi

# Test Refresh Token
echo ""
echo "16. Testing Refresh Token..."
REFRESH=$(curl -s -X POST "$BASE_URL/auth/refresh" \
    -H "Content-Type: application/json" \
    -d "{\"refresh_token\":\"$REFRESH_TOKEN\"}")
if echo "$REFRESH" | grep -q '"access_token"'; then
    echo "   ✓ Refresh Token passed"
else
    echo "   ✗ Refresh Token failed: $REFRESH"
fi

# Test Logout
echo ""
echo "17. Testing Logout..."
LOGOUT=$(curl -s -X POST "$BASE_URL/auth/logout" -H "$AUTH")
if echo "$LOGOUT" | grep -q '"success":true'; then
    echo "   ✓ Logout passed"
else
    echo "   ✗ Logout failed: $LOGOUT"
fi

echo ""
echo "============================================"
echo "Testing completed!"
echo "============================================"
