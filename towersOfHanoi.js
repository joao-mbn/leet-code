function main() {
  const spikes = {
    from: [11, 9, 7, 5, 3],
    to: [],
    spare: [],
  }
  const nDisks = spikes.from.length

  move(nDisks, "from", "to", "spare");

  function prettyPrint() {
    const maxDiskSize = Math.max(...spikes.from, ...spikes.to, ...spikes.spare, 0);
    const space = maxDiskSize + 2;

    const centerDiskInSpace = (disk) => {
      const diskSize = disk || 0;
      const spaceLeft = Math.floor((space - diskSize) / 2);
      const spaceRight = space - diskSize - spaceLeft;

      return " ".repeat(spaceLeft) + "#".repeat(disk) + " ".repeat(spaceRight);
    }

    for (let i = nDisks - 1; i >= 0; i--) {
      const fromTop = spikes.from[i] || 0;
      const toTop = spikes.to[i] || 0;
      const spareTop = spikes.spare[i] || 0;

      console.log(centerDiskInSpace(fromTop) + centerDiskInSpace(toTop) + centerDiskInSpace(spareTop));
    }

    console.log('\n', "-".repeat(space * 3), '\n');
  }

  /**
   * @param {number[]} from
   * @param {number[]} to
   */
  function singleMove(from, to) {
    const top = spikes[from].pop()
    spikes[to].push(top)

    prettyPrint()
  }

  function move(n, from, to, spare) {
    if (n === 1) {
      singleMove(from, spare)
      singleMove(from, to)
      singleMove(spare, to)

      return
    }

    move(n - 1, from, spare, to)

    if (n === nDisks) return

    singleMove(from, to)
    move(n - 1, spare, to, from)
  }
}

main()