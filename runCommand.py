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
    runCommand('cd C:\\Riot Games\\League of Legends\\RADS\\solutions\\lol_game_client_sln\\releases\\0.0.1.198 & start "" /D "deploy" "League of Legends.exe" "8394" "LoLLauncher.exe" "" "replay spectator.br1.lol.riotgames.com:80 WN2SG+pd3mNBnzh7xsXUD9v/PUyAALsO 1261212781 BR1"')

if __name__ == "__main__":
    main()
