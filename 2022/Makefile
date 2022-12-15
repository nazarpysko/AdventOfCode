new:
	@echo "Creating new structure for day $(day) ..."

	@mkdir day$(day)
	@echo "package main" > ./day$(day)/common.go
	@cp template ./day$(day)/part1.go | sed -i -e 's/FuncName/Part1/' -e 's/day/$(day)/' ./day$(day)/part1.go
	@cp template ./day$(day)/part2.go | sed -i -e 's/FuncName/Part2/' -e 's/day/$(day)/' ./day$(day)/part2.go

	@echo "Done!"

copyPart1:
	@echo "From day $(day), copying contents of part1 to part2..."
	@cp ./day$(day)/part1.go ./day$(day)/part2.go | sed -i 's/Part1/Part2/' ./day$(day)/part2.go
	@echo "Done!"