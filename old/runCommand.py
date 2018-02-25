import os
import subprocess

def runCommand(cmd):
    result = []
    process = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE,stderr=subprocess.PIPE)
    for line in process.stdout:
        result.append(line)
    for line in result:
        print(line)
    print(process.stderr)
    errCode = process.returncode
    if errCode is not None:
        raise Exception('cmd %s failed, see above for details', cmd)

def main():
    print(os.getcwd())
    # cd "c:\Riot Games\League of Legends\RADS\solutions\lol_game_client_sln\releases\" & for /d %F in (*) do cd %F & start "" /D "deploy" "League of Legends.exe" "8394" "LoLLauncher.exe" "" "replay spectator.br1.lol.riotgames.com:80 OKWIk9sFl1zrJZcWXs52C+DKH+wQIuf0 1261583214 BR1"
    runCommand('cd C:\\Riot Games\\League of Legends\\RADS\\solutions\\lol_game_client_sln\\releases\\0.0.1.198" start "" "deploy" "League of Legends.exe" "8394" "LoLLauncher.exe" "" "spectator spectator.br1.lol.riotgames.com:80 Inf/3Xp6+8aXgDsKnxRYCwbtI9Me0lDZ 1261585128 BR1" "-UseRads"')

if __name__ == "__main__":
    main()
