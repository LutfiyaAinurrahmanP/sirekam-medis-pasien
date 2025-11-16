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

# Hitung test suites (file test)
$testFiles = @("medicine_test.go", "test_type_test.go")
$passedSuites = 0
$failedSuites = 0

foreach ($file in $testFiles) {
    if ($testOutput -match "PASS.*$file" -or ($testOutput -match "$file" -and $exitCode -eq 0)) {
        $passedSuites++
    } elseif ($testOutput -match "FAIL.*$file") {
        $failedSuites++
    }
}

$totalSuites = $testFiles.Count

# Clear screen dan tampilkan hasil
Clear-Host
Write-Host ""

# Status badge untuk setiap file
foreach ($file in $testFiles) {
    if ($testOutput -match "FAIL.*$file") {
        Write-Host " FAIL " -BackgroundColor Red -ForegroundColor Black -NoNewline
        Write-Host "  tests/$file" -ForegroundColor Red
    } else {
        Write-Host " PASS " -BackgroundColor Green -ForegroundColor Black -NoNewline
        Write-Host "  tests/$file" -ForegroundColor Green
    }
}

Write-Host ""

# Summary
Write-Host "Test Suites: " -NoNewline
if ($failedSuites -gt 0) {
    Write-Host "$failedSuites failed, " -NoNewline -ForegroundColor Red
}
if ($passedSuites -gt 0) {
    Write-Host "$passedSuites passed" -NoNewline -ForegroundColor Green
    Write-Host ", " -NoNewline
}
Write-Host "$totalSuites total"

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
