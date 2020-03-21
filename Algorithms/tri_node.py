"""
参考：
  https://rightcode.co.jp/blog/information-technology/trie-fast-dictionary-implementation-2
"""

VOCAB_SIZE = 128


class TrieNode:
  """
  TriTree's one node class

  Attributes
  ----------
  item : str
    registered word's yomigana

  children : list
    list of next node
  """

  def __init__(self, item=None):
    self.item = item
    self.children = [-1 for i in range(VOCAB_SIZE)]

  def __str__(self):
    ret = ''
    ret = ret + 'item = {}\n'.format(self.item)
    ret = ret + 'children = {}\n'.format(self.children)
    return ret


class Trie:
  """
  Trie Tree class

  Attributes
  ----------
  nodes : TrieNode
    node's list
  """

  def __init__(self):
    root = TrieNode()
    self.nodes = [root]

  def __add_node(self, node):
    """
    add node to nodes
    return added node's index

    Parameters
    ----------
    node : TrieNode
      add node

    Returns
    -------
    index : int
      added node's index
    """
    self.nodes.append(node)
    return len(self.nodes) - 1

  def __get_char_num(self, c):
    """
    return c's id
    ex: a->1, b->2, c->3, ..., z->26

    Parameters
    ----------
    c : char

    Returns
    -------
    id : int
      c's id
    """
    return ord(c) - ord('a')

  def insert(self, word, item, char_index=0, node_index=0):
    """
    register new word to Trie

    Parameters
    ----------
    word : str
      word which u wanna register

    item : str
      word's yomigana

    char_index : int
     focused position of chars

    node_index : int
      focused position of nodes
    """

    char_num = self.__get_char_num(word[char_index])
    next_node_index = self.nodes[node_index].children[char_num]
    if next_node_index == -1:
      new_node = TrieNode()
      next_node_index = self.__add_node(new_node)
      self.nodes[node_index].children[char_num] = next_node_index

    if char_index == len(word) - 1:
      self.nodes[next_node_index].item = item
    else:
      self.insert(word, item, char_index+1, next_node_index)

  def query(self, word):
    """
      search Trie

      Parameters
      ----------
      word : str
        search word

      Returns
      -------
      item : str or None
        search result
    """
    node_index = 0
    for c in word:
      char_num = self.__get_char_num(c)
      tmp_node = self.nodes[node_index]

      node_index = tmp_node.children[char_num]
      if node_index == -1:
        return None

    return self.nodes[node_index].item



# スクリプト実行
trie = Trie()

trie.insert("ab", "エービー")
trie.insert("abc", "エービーシー")
trie.insert("b", "ビー")
trie.insert("bc", "ビーシー")
trie.insert("c", "シー")

print(trie.query("a"))
print(trie.query("abc"))
print(trie.query("bc"))
print(trie.query("abcd"))