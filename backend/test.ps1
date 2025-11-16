# Script sederhana untuk menampilkan hasil test dengan format Jest-like

# Jalankan test dan hitung waktu
$startTime = Get-Date
$testOutput = go test ./tests/... -v 2>&1 | Out-String
$exitCode = $LASTEXITCODE
$endTime = Get-Date
$duration = ($endTime - $startTime).TotalSeconds

# Hitung test cases
$passedTests = ($testOutput | Select-String -Pattern "--- PASS:" -AllMatches).Matches.Count
$failedTests = ($testOutput | Select-String -Pattern "--- FAIL:" -AllMatches).Matches.Count
$totalTests = $passedTests + $failedTests

# Clear screen dan tampilkan hasil
Clear-Host
Write-Host ""

# Status badge
if ($exitCode -eq 0) {
    Write-Host " PASS " -BackgroundColor Green -ForegroundColor Black -NoNewline
    Write-Host "  tests/medicine_test.go" -ForegroundColor Green
} else {
    Write-Host " FAIL " -BackgroundColor Red -ForegroundColor Black -NoNewline
    Write-Host "  tests/medicine_test.go" -ForegroundColor Red
}

Write-Host ""

# Summary
Write-Host "Test Suites: " -NoNewline
if ($exitCode -eq 0) {
    Write-Host "1 passed" -NoNewline -ForegroundColor Green
} else {
    Write-Host "1 failed" -NoNewline -ForegroundColor Red
}
Write-Host ", 1 total"

Write-Host "Tests:       " -NoNewline
if ($failedTests -gt 0) {
    Write-Host "$failedTests failed, " -NoNewline -ForegroundColor Red
}
Write-Host "$passedTests passed" -NoNewline -ForegroundColor Green
Write-Host ", $totalTests total"

Write-Host "Snapshots:   0 total"
Write-Host ("Time:        {0:F3} s" -f $duration)
Write-Host ""
Write-Host "Ran all test suites." -ForegroundColor DarkGray
Write-Host ""

exit $exitCode
