# AutoRPG - self playing adventure! (not really)

![Pipeline](https://github.com/RafaelRochaS/autorpg/actions/workflows/main.yml/badge.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=RafaelRochaS_autorpg&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=RafaelRochaS_autorpg)
[![Go Report Card](https://goreportcard.com/badge/github.com/RafaelRochaS/autorpg)](https://goreportcard.com/report/github.com/RafaelRochaS/autorpg)

AutoRPG is an evolution of a previous project, a simple text-based RPG, where you fight randomnly generated enemies, level up, and gain new equipment. However, unlike the previous project, this game plays out all by itself. You simply pick a name, class, base stats, and the game fights random enemies until you die. Its not really meant for fun. But it is a fun software engineering experiment!


## Tech Stack

Pure go, only. A couple of libs for some basic stuff, and that is that.


## Settings

The system allows for a debug mode to be set, where all the basic choices are done automatically, and a lot of information about the behind-the-scenes is printed. 

You can also set the combat time (in seconds), which controls how long the game waits between combats. If set to 0, you can blow through a lot of enemies in a very short amount of time. Until you die, that is.

These can be set through environment variables, such as:


    DEBUG=False 
    COMBAT_PAUSE_TIME=0

If not set, the combat time will default to 1.


## Finally

That is basically it! All the information about classes and stuff and printed when you first run the program. Oh, and everything is extremelly imbalanced at this point, because it is very hard to balance games (even fake ones).
